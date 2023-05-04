package handlers

import (
	"os"
	"testing"

	"github.com/Rha02/resumanager/src/dbrepo"
	authtokenservice "github.com/Rha02/resumanager/src/services/authTokenService"
	cacheservice "github.com/Rha02/resumanager/src/services/cacheService"
	filestorageservice "github.com/Rha02/resumanager/src/services/fileStorageService"
	"github.com/go-chi/chi/v5"
)

// TestMain is the entry point for all tests in this package.
func TestMain(m *testing.M) {
	// init repos
	dbRepo := dbrepo.NewTestDBRepo()
	cacheRepo := cacheservice.NewTestRepo()
	fileStorageRepo := filestorageservice.NewTestFileStorage()
	authTokenRepo := authtokenservice.NewTestAuthTokenRepo()

	// init cacheRepo data
	cacheRepo.Set("blacklisted_token", "true", 0)

	// init handlers
	NewHandlers(NewRepository(
		dbRepo,
		cacheRepo,
		fileStorageRepo,
		authTokenRepo,
	))

	os.Exit(m.Run())
}

func getRoutes() *chi.Mux {
	mux := chi.NewRouter()

	mux.Post("/login", Repo.Login)
	mux.Post("/refresh", Repo.Refresh)
	mux.Post("/register", Repo.Register)
	mux.Post("/logout", Repo.Logout)

	return mux
}
