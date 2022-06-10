package ma

import (
	"log"
	"net/http"

	"github.com/pittsCourt/Server2/handlers"
)

func main() {
	// Handling the /data/1 as a function
	http.HandleFunc("/data/1", handlers.Handler)

	log.Println("Listening on port :80")

	http.ListenAndServe(":80", nil)
}
