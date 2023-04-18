package main

import (
	"log"
	"net/http"

	"github.com/Rha02/resumanager/src/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

var PORT = ":3000"

func main() {
	godotenv.Load()
	// init handlers
	handlers.NewHandlers(handlers.NewRepository())

	router := newRouter()

	log.Printf("Server is running on port %s", PORT)
	http.ListenAndServe(PORT, router)
}

func newRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Post("/login", handlers.Repo.Login)
	r.Post("/register", handlers.Repo.Register)
	r.Post("/logout", handlers.Repo.Logout)
	r.Post("/refresh", handlers.Repo.Refresh)
	r.Get("/checkauth", handlers.Repo.CheckAuth)

	return r
}
