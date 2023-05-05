package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Rha02/resumanager/src/dbrepo"
	"github.com/Rha02/resumanager/src/driver"
	"github.com/Rha02/resumanager/src/handlers"
	"github.com/Rha02/resumanager/src/middleware"
	authtokenservice "github.com/Rha02/resumanager/src/services/authTokenService"
	cacheservice "github.com/Rha02/resumanager/src/services/cacheService"
	filestorageservice "github.com/Rha02/resumanager/src/services/fileStorageService"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

var PORT = ":3000"

func main() {
	godotenv.Load()

	db, err := driver.ConnectSQL(os.Getenv("DB_CONNECTION"))
	if err != nil {
		log.Fatal("Cannot connect to database! Dying...")
	}

	defer db.Close()

	// init db repo
	dbRepo := dbrepo.NewPostgresRepo(db.SQL)

	// init blacklist for auth refresh tokens
	blacklist := cacheservice.NewRedisRepo(os.Getenv("REDIS_ADDRESS"), os.Getenv("REDIS_PASSWORD"))
	defer blacklist.Close()

	// init file storage service
	fileStorageRepo := filestorageservice.NewAzureFileStorage(
		os.Getenv("AZURE_STORAGE_ACCOUNT_NAME"),
		os.Getenv("AZURE_STORAGE_ACCOUNT_KEY"),
	)

	// init auth token service
	authTokenRepo := authtokenservice.NewAuthTokenProvider(os.Getenv("JWT_SIGNING_ALGORITHM"))

	// init handlers
	handlers.NewHandlers(handlers.NewRepository(
		dbRepo, blacklist, fileStorageRepo, authTokenRepo,
	))

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

		r.Get("/resumes", handlers.Repo.GetUserResumes)
		r.Get("/resumes/{resumeID}", handlers.Repo.GetResume)
		r.Post("/resumes", handlers.Repo.PostResume)
		r.Delete("/resumes/{resumeID}", handlers.Repo.DeleteResume)
	})

	return r
}
