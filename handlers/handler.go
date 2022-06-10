package handlers

import (
	"net/http"
)

func Handler(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json; charset=utf-8")
	res.WriteHeader(200)
	res.Write(b)
}
