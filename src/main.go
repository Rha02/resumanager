package main

import (
	"log"
	"net/http"

	"github.com/Rha02/resumanager/src/handlers"
	"github.com/go-chi/chi/v5"
)

var PORT = ":3000"

func main() {
	// init handlers
	handlers.NewHandlers(handlers.NewRepository())

	router := newRouter()

	log.Printf("Server is running on port %s", PORT)
	http.ListenAndServe(PORT, router)
}

func newRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/", handlers.Repo.Home)

	return r
}
