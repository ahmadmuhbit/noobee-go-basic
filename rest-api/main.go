package main

import (
	"log"
	"net/http"
	"rest-api/menu"

	"github.com/go-chi/chi/v5"
)

func main() {
	router := chi.NewRouter()

	menu.Register(router, nil)
	const port = ":8000"

	log.Println("Server running at port", port)
	http.ListenAndServe(port, router)
}
