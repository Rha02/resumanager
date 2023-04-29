package dbrepo

import "github.com/Rha02/resumanager/src/models"

// DatabaseRepository is an interface for database operations
type DatabaseRepository interface {
	GetUserByID(id string) (models.User, error)
	GetUserByUsername(username string) (models.User, error)
	CreateUser(user models.User) (int, error)
}
