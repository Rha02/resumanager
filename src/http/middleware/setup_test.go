package middleware

import (
	"net/http"
	"os"
	"testing"

	"github.com/Rha02/resumanager/src/dbrepo"
	"github.com/Rha02/resumanager/src/http/handlers"
	authtokenservice "github.com/Rha02/resumanager/src/services/authTokenService"
	cacheservice "github.com/Rha02/resumanager/src/services/cacheService"
	filestorageservice "github.com/Rha02/resumanager/src/services/fileStorageService"
	hashservice "github.com/Rha02/resumanager/src/services/hashService"
	"github.com/go-chi/chi/v5"
)

// globals
var handler *chi.Mux

func TestMain(m *testing.M) {
	// init repos
	dbRepo := dbrepo.NewTestDBRepo()
	cacheRepo := cacheservice.NewTestRepo()
	fileStorageRepo := filestorageservice.NewTestFileStorage()
	authTokenRepo := authtokenservice.NewTestAuthTokenRepo()
	hashRepo := hashservice.NewTestHashRepo()

	// init handlers
	handlers.NewHandlers(handlers.NewRepository(
		dbRepo,
		cacheRepo,
		fileStorageRepo,
		authTokenRepo,
		hashRepo,
	))

	handler = getRoutes()

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
