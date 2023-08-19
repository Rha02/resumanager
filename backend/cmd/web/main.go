package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Rha02/resumanager/src/dbrepo"
	"github.com/Rha02/resumanager/src/driver"
	"github.com/Rha02/resumanager/src/http/handlers"
	"github.com/Rha02/resumanager/src/http/middleware"
	authtokenservice "github.com/Rha02/resumanager/src/services/authTokenService"
	cacheservice "github.com/Rha02/resumanager/src/services/cacheService"
	filestorageservice "github.com/Rha02/resumanager/src/services/fileStorageService"
	hashservice "github.com/Rha02/resumanager/src/services/hashService"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

var PORT = ":80"

func main() {
	godotenv.Load()

	db, err := driver.ConnectSQL(os.Getenv("DB_CONNECTION"))
	if err != nil {
		log.Fatal("Cannot connect to database! Dying...")
	}
	defer db.Close()

	rdb, err := driver.ConnectRedis(os.Getenv("REDIS_ADDRESS"), os.Getenv("REDIS_PASSWORD"))
	if err != nil {
		log.Fatal("Cannot connect to redis! Dying...")
	}
	defer rdb.Close()

	// init db repo
	dbRepo := dbrepo.NewPostgresRepo(db.SQL)

	// init blacklist for auth refresh tokens
	blacklist := cacheservice.NewRedisRepo(rdb.Redis)

	// init file storage service
	fileStorageRepo := filestorageservice.NewAzureFileStorage(
		os.Getenv("AZURE_STORAGE_ACCOUNT_NAME"),
		os.Getenv("AZURE_STORAGE_ACCOUNT_KEY"),
	)

	// init auth token service
	authTokenRepo := authtokenservice.NewAuthTokenProvider(os.Getenv("JWT_SIGNING_ALGORITHM"))

	// init hash service
	hashRepo := hashservice.NewBcryptRepo()

	// init handlers
	handlers.NewHandlers(handlers.NewRepository(
		dbRepo, blacklist, fileStorageRepo, authTokenRepo, hashRepo,
	))

	router := newRouter()

	server := &http.Server{
		Addr:    PORT,
		Handler: router,
	}

	log.Printf("Server is running on port %s", PORT)

	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
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
