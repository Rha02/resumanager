package dbrepo

import "github.com/Rha02/resumanager/src/models"

// DatabaseRepository is an interface for database operations
type DatabaseRepository interface {
	GetUserByID(id string) (models.User, error)
	GetUserByEmail(username string) (models.User, error)
	CreateUser(user models.User) (string, error)

	GetUserResumes(userID string) ([]models.Resume, error)
	GetResume(id string) (models.Resume, error)
	InsertResume(resume models.Resume) error
	DeleteResume(id string) error
}
