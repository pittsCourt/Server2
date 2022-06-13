package main

import (
	"log"
	"net/http"

	"github.com/digicert/health"
	"github.com/pittsCourt/Server2/handlers"
)

func main() {
	health.SetLogLevel("debug")
	health.SetDebug(true)

	// Handling the /data/ paths
	http.HandleFunc("/data/", handlers.DataHandler)

	// Handling the /data/2 path
	// http.HandleFunc("/data/2", handlers.DataHandler2)

	log.Println("Listening on port :80")

	http.ListenAndServe(":80", nil)
}
