package health

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/prometheus/client_golang/prometheus"
)

/*
 * Collectors. Things that sit in the middle of things or get called to collect data
 */

/*
 *	Register a collector defined in another module
 */
func Register(collector prometheus.Collector) error {
	return prometheus.Register(collector)
}

func LoggingHTTPClient() *http.Client {
	return loggingHTTPClient
}

/*
 *	loggingMiddleware sits in the middle between a function and the user, and will run (and therefore log) whenever a user calls the router
 *   @todo the request email isn't working here, and only in the report.go hanlder functions. Fix this so we get an email tag on every hit
 */
 func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		path := strings.Split(request.RequestURI, "?")[0]

		timer := prometheus.NewTimer(httpDuration.WithLabelValues(path))
		defer timer.ObserveDuration()
		rw := NewResponseWriter(response)
		next.ServeHTTP(rw, request)

		statusCode := rw.statusCode

		responseStatus.WithLabelValues(strconv.Itoa(statusCode), "inbound").Inc()
		totalRequests.WithLabelValues(parsePath(path), "inbound").Inc()

		Trace("%v %v %v %v", rw.requestorEmail, request.RemoteAddr, request.Method, request.RequestURI)
	})
}

type LoggingRoundTripper struct {
	Proxied http.RoundTripper
}

/*
 * RoundTrip replaces a section of the http.Client that allows it to log and add metrics
 * 	for activity on that client
 */
func (lrt LoggingRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	Trace("Sending request to %v", req.URL.Path)
	timer := prometheus.NewTimer(httpDuration.WithLabelValues(req.URL.Path))
	defer timer.ObserveDuration()

	res, err := lrt.Proxied.RoundTrip(req)
	if err != nil {
		Error("Encountered error getting round trip request %v", err)
		responseStatus.WithLabelValues("500", "outbound").Inc()
		totalRequests.WithLabelValues(parsePath(req.URL.Path), "outbound").Inc()
		return res, err
	}

	Trace("Received %v response", res.Status)
	path := strings.Split(req.RequestURI, "?")[0]
	responseStatus.WithLabelValues(strconv.Itoa(res.StatusCode), "outbound").Inc()
	totalRequests.WithLabelValues(parsePath(path), "outbound").Inc()
	return res, err
}

// Wrapping the http.ResponseWriter to make an upstream trackable statusCode value
type responseWriter struct {
	http.ResponseWriter
	statusCode     int
	requestorEmail string
}

func NewResponseWriter(w http.ResponseWriter) *responseWriter {
	return &responseWriter{w, http.StatusOK, ""}
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

func (rw *responseWriter) RequesterEmail(email string) {
	rw.requestorEmail = email
}