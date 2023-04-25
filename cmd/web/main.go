package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Rha02/resumanager/src/handlers"
	"github.com/Rha02/resumanager/src/middleware"
	authtokenservice "github.com/Rha02/resumanager/src/services/authTokenService"
	cacheservice "github.com/Rha02/resumanager/src/services/cacheService"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

var PORT = ":3000"

func main() {
	godotenv.Load()

	// init cache
	cacheRepo := cacheservice.NewTestCacheRepo()

	// init auth token service
	authTokenRepo := authtokenservice.NewAuthTokenProvider(os.Getenv("JWT_SIGNING_ALGORITHM"))

	// init handlers
	handlers.NewHandlers(handlers.NewRepository(cacheRepo, authTokenRepo))

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

	r.Group(func(r chi.Router) {
		r.Use(middleware.RequiresAuthentication(handlers.Repo))
		r.Get("/checkauth", handlers.Repo.CheckAuth)
	})

	return r
}
