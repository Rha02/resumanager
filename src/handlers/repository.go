package handlers

import (
	"github.com/Rha02/resumanager/src/dbrepo"
	authtokenservice "github.com/Rha02/resumanager/src/services/authTokenService"
	cacheservice "github.com/Rha02/resumanager/src/services/cacheService"
	"github.com/go-playground/validator/v10"
)

type ContextKey struct{}

type Repository struct {
	DB            dbrepo.DatabaseRepository
	Blacklist     cacheservice.CacheRepository
	AuthTokenRepo authtokenservice.AuthTokenRepository
	validator     validator.Validate
}

var Repo *Repository

func NewRepository(
	db dbrepo.DatabaseRepository,
	cacheRepo cacheservice.CacheRepository,
	authTokenRepo authtokenservice.AuthTokenRepository,
) *Repository {
	return &Repository{
		DB:            db,
		Blacklist:     cacheRepo,
		AuthTokenRepo: authTokenRepo,
		validator:     *validator.New(),
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}
