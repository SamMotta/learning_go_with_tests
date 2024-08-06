package main

import (
	"net/http"

	"github.com/SamMotta/go-specs-greet/adapters/httpserver"
)

func main() {
	mux := httpserver.NewHandler()
	http.ListenAndServe(":8080", mux)
}
