package handlers

import (
	"os"
	"testing"

	cacheservice "github.com/Rha02/resumanager/src/services/cacheService"
	"github.com/go-chi/chi/v5"
)

// TestMain is the entry point for all tests in this package.
func TestMain(m *testing.M) {
	cacheRepo := cacheservice.NewTestCacheRepo()

	// init handlers
	NewHandlers(NewRepository(cacheRepo))

	os.Exit(m.Run())
}

func getRoutes() *chi.Mux {
	mux := chi.NewRouter()

	mux.Post("/login", Repo.Login)

	return mux
}
