package handlers

import (
	authtokenservice "github.com/Rha02/resumanager/src/services/authTokenService"
	cacheservice "github.com/Rha02/resumanager/src/services/cacheService"
)

type ContextKey struct{}

type Repository struct {
	CacheRepo     cacheservice.CacheRepository
	AuthTokenRepo authtokenservice.AuthTokenRepository
}

var Repo *Repository

func NewRepository(cacheRepo cacheservice.CacheRepository) *Repository {
	return &Repository{
		CacheRepo: cacheRepo,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}
