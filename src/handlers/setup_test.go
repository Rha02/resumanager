package handlers

import (
	"os"
	"testing"

	"github.com/Rha02/resumanager/src/dbrepo"
	"github.com/Rha02/resumanager/src/models"
	authtokenservice "github.com/Rha02/resumanager/src/services/authTokenService"
	cacheservice "github.com/Rha02/resumanager/src/services/cacheService"
	"github.com/go-chi/chi/v5"
)

// TestMain is the entry point for all tests in this package.
func TestMain(m *testing.M) {
	// init repos
	dbRepo := dbrepo.NewTestDBRepo()
	cacheRepo := cacheservice.NewTestCacheRepo()
	authTokenRepo := authtokenservice.NewTestAuthTokenRepo()

	// init dbRepo data
	dbRepo.CreateUser(models.User{
		Username: "testuser",
		Password: "testpassword",
	})
	dbRepo.CreateUser(models.User{
		Username: "access_token_error",
		Password: "testpassword",
	})
	dbRepo.CreateUser(models.User{
		Username: "refresh_token_error",
		Password: "testpassword",
	})

	// init cacheRepo data
	cacheRepo.Set("blacklisted_token", "true")

	// init handlers
	NewHandlers(NewRepository(
		dbRepo,
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
