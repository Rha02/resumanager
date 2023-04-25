package middleware

import (
	"net/http"
	"os"
	"testing"

	"github.com/Rha02/resumanager/src/handlers"
	authtokenservice "github.com/Rha02/resumanager/src/services/authTokenService"
	cacheservice "github.com/Rha02/resumanager/src/services/cacheService"
	"github.com/go-chi/chi/v5"
)

func TestMain(m *testing.M) {
	cacheRepo := cacheservice.NewTestCacheRepo()
	authTokenRepo := authtokenservice.NewTestAuthTokenRepo()

	// init handlers
	handlers.NewHandlers(handlers.NewRepository(
		cacheRepo,
		authTokenRepo,
	))

	os.Exit(m.Run())
}

func getRoutes() *chi.Mux {
	mux := chi.NewRouter()

	mux.Group(func(r chi.Router) {
		r.Use(RequiresAuthentication(handlers.Repo))
		r.Post("/test-middleware/auth", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("OK"))
		})
	})

	return mux
}