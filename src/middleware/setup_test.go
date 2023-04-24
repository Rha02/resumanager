package middleware

import (
	"net/http"
	"os"
	"testing"

	authtokenservice "github.com/Rha02/resumanager/src/services/authTokenService"
	"github.com/go-chi/chi/v5"
)

func TestMain(m *testing.M) {
	// init AuthTokenService
	authtokenservice.NewAuthTokenRepo(authtokenservice.NewTestAuthTokenRepo())

	os.Exit(m.Run())
}

func getRoutes() *chi.Mux {
	mux := chi.NewRouter()

	mux.Group(func(r chi.Router) {
		r.Use(RequiresAuthentication)
		r.Post("/test-middleware/auth", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("OK"))
		})
	})

	return mux
}
