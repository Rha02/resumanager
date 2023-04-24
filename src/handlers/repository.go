package handlers

import (
	cacheservice "github.com/Rha02/resumanager/src/services/cacheService"
)

type Repository struct {
	CacheRepo cacheservice.CacheRepository
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
