package handlers

import (
	"os"
	"testing"

	authtokenservice "github.com/Rha02/resumanager/src/services/authTokenService"
	cacheservice "github.com/Rha02/resumanager/src/services/cacheService"
	"github.com/go-chi/chi/v5"
)

// TestMain is the entry point for all tests in this package.
func TestMain(m *testing.M) {
	cacheRepo := cacheservice.NewTestCacheRepo()
	authTokenRepo := authtokenservice.NewTestAuthTokenRepo()

	// init handlers
	NewHandlers(NewRepository(
		cacheRepo,
		authTokenRepo,
	))

	os.Exit(m.Run())
}

func getRoutes() *chi.Mux {
	mux := chi.NewRouter()

	mux.Post("/login", Repo.Login)
	mux.Post("/refresh", Repo.Refresh)
	mux.Post("/logout", Repo.Logout)

	return mux
}
