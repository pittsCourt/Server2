package ma

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Data struct {
	Id    int    //`json:"id"`
	Value string //`json:"value"`
}

func DataHandler(w http.ResponseWriter, r *http.Request) {
	m := Data{1, "one"}
	b, _ := json.Marshal(m)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(200)
	w.Write(b)
	data := r.URL.Path[len("/data/"):]
	fmt.Println(data)
}

func main() {
	// Handling the /data/1 as a function
	http.HandleFunc("/data/1", DataHandler)

	log.Println("Listening on port :80")

	http.ListenAndServe(":80", nil)
}
