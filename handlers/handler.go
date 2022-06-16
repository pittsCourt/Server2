package handlers

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/digicert/health"
)

// func configVariables() {
// 	configFile, err := ioutil.ReadFile("config.yaml")
// 	if err != nil {

// 		log.Fatal(err)
// 	}

// 	err2 := yaml.Unmarshal(configFile)

// 	if err2 != nil {

// 		log.Fatal(err2)
// 	}

// 	health.Debug("%s", err2)

// }

func getData(sFull string) []byte {
	// Get response from Server found at sFull
	resp, err := http.Get(sFull)
	if err != nil {
		health.Fatal("This is the error: %v", err)
	}

	// Gather response from Server and return it to DataHandler
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	return body
}

func DataHandler(w http.ResponseWriter, r *http.Request) {
	// Declare s variable to the url path after /data/
	s := r.URL.Path[len("/data/"):]
	if s == "" {
		health.Error("Path after '/data/' is empty.")
		w.WriteHeader(500)
		w.Write([]byte("No data to display, try a number after '/data/' in the URL"))
		return
	}
	health.Debug("Path contains a value %s", s)
	// Convert s to string
	sStr := string(s)
	health.Debug("%s is now a string", sStr)
	// Append sSTring to localhost url
	sFull := "http://server1:8080/data/" + sStr
	// sFull := "http://localhost:8080/data/" + sStr

	// Call get data function with path to data
	body := getData(sFull)

	// Write data to page
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(200)
	w.Write(body)
}

/*
No longer used, but here for reference

func DataHandler2(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("http://localhost:8080/data/2")
	if err != nil {
		log.Fatalln(err)
	}

	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	// Write data to page
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(200)
	w.Write(body)
}
*/
