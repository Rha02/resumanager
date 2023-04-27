package handlers

import (
	authtokenservice "github.com/Rha02/resumanager/src/services/authTokenService"
	cacheservice "github.com/Rha02/resumanager/src/services/cacheService"
)

type ContextKey struct{}

type Repository struct {
	Blacklist     cacheservice.CacheRepository
	AuthTokenRepo authtokenservice.AuthTokenRepository
}

var Repo *Repository

func NewRepository(cacheRepo cacheservice.CacheRepository, authTokenRepo authtokenservice.AuthTokenRepository) *Repository {
	return &Repository{
		Blacklist:     cacheRepo,
		AuthTokenRepo: authTokenRepo,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}
