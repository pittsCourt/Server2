package health

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	debug             bool
	logLevel          int
	loggingHTTPClient *http.Client
)

func init() {
	debug = false
	SetLogLevel("info")

	loggingHTTPClient = &http.Client{
		Transport: LoggingRoundTripper{http.DefaultTransport},
	}
}

/*
 * The common metrics. I.e. Things that get tracked
 */
var responseStatus = promauto.NewCounterVec(
	prometheus.CounterOpts{
		Name: "response_code",
		Help: "Status of HTTP responses",
	},
	[]string{"status", "direction"},
)

var totalRequests = promauto.NewCounterVec(
	prometheus.CounterOpts{
		Name: "total_requests",
		Help: "Total number of requests",
	},
	[]string{"path", "direction"},
)

var httpDuration = promauto.NewHistogramVec(
	prometheus.HistogramOpts{
		Name: "http_average_time_seconds",
		Help: "Duration of HTTP requests.",
	},
	[]string{"path"},
)

var logCount = promauto.NewCounterVec(
	prometheus.CounterOpts{
		Name: "log_count",
		Help: "Count of log entries, by type",
	},
	[]string{"type"},
)