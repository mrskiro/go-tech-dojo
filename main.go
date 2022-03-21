package main

import (
	"log"
	"net/http"

	h "github.com/purp1eeeee/go-tech-dojo/handlers/http"
	"github.com/purp1eeeee/go-tech-dojo/oapi"
)

func main() {
	handlers := h.NewHandlers()
	handler := oapi.Handler(handlers)

	mux := http.NewServeMux()
	mux.Handle("/", handler)
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}
