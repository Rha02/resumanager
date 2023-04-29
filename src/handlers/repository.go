package handlers

import (
	"github.com/Rha02/resumanager/src/dbrepo"
	authtokenservice "github.com/Rha02/resumanager/src/services/authTokenService"
	cacheservice "github.com/Rha02/resumanager/src/services/cacheService"
)

type ContextKey struct{}

type Repository struct {
	DB            dbrepo.DatabaseRepository
	Blacklist     cacheservice.CacheRepository
	AuthTokenRepo authtokenservice.AuthTokenRepository
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
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}
