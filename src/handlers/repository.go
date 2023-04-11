package handlers

type Repository struct{}

var Repo *Repository

func NewRepository() *Repository {
	return &Repository{}
}

func NewHandlers(r *Repository) {
	Repo = r
}
