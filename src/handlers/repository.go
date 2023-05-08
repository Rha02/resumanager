package handlers

import (
	"github.com/Rha02/resumanager/src/dbrepo"
	authtokenservice "github.com/Rha02/resumanager/src/services/authTokenService"
	cacheservice "github.com/Rha02/resumanager/src/services/cacheService"
	filestorageservice "github.com/Rha02/resumanager/src/services/fileStorageService"
	hashservice "github.com/Rha02/resumanager/src/services/hashService"
	"github.com/go-playground/validator/v10"
)

type ContextKey struct{}

type Repository struct {
	DB            dbrepo.DatabaseRepository
	Blacklist     cacheservice.CacheRepository
	FileStorage   filestorageservice.FileStorageRepository
	AuthTokenRepo authtokenservice.AuthTokenRepository
	hashRepo      hashservice.HashRepository
	validator     validator.Validate
}

var Repo *Repository

func NewRepository(
	db dbrepo.DatabaseRepository,
	cacheRepo cacheservice.CacheRepository,
	fileStorage filestorageservice.FileStorageRepository,
	authTokenRepo authtokenservice.AuthTokenRepository,
) *Repository {
	return &Repository{
		DB:            db,
		Blacklist:     cacheRepo,
		FileStorage:   fileStorage,
		AuthTokenRepo: authTokenRepo,
		hashRepo:      hashservice.NewBcryptRepo(),
		validator:     *validator.New(),
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}
