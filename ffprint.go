package main

import (
	"net/http"
	"github.com/go-chi/chi"
	"go-training/handler"
)

func main() {
	r := chi.NewRouter()

	// define endpoint and associate with handler
	r.Get("/list", handler.ListHandler)

	// setup http server on port 8080
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}
