package main

import (
	"log"
	"net/http"
)

func DataHandler1(w http.ResponseWriter, r *http.Request) {
	// data := r.URL.Path[len("/data/"):]
	// if data == "1" {

	// } else {

	// }

	// data from ":8080/data/1"
	// b :=
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(200)
	// w.Write(b)
}

func main() {
	// Handling the /data/1 as a function
	http.HandleFunc("/data/1", DataHandler)

	log.Println("Listening on port :80")

	http.ListenAndServe(":80", nil)
}
