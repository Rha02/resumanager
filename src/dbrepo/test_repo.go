package dbrepo

import (
	"errors"

	"github.com/Rha02/resumanager/src/models"
)

type testDBRepo struct{}

// NewTestDBRepo creates a new test database repository
func NewTestDBRepo() DatabaseRepository {
	return &testDBRepo{}
}

// GetUserByID gets a user by ID
func (m *testDBRepo) GetUserByID(id string) (models.User, error) {
	var user models.User
	if id == "db_get_user_error" {
		return user, errors.New("error getting user")
	}

	user.ID = 1
	user.Username = "testuser"

	return user, nil
}

// GetUserByUsername gets a user by username
func (m *testDBRepo) GetUserByUsername(username string) (models.User, error) {
	var user models.User
	if username == "db_get_user_error" {
		return user, errors.New("error getting user")
	}

	user.Username = username
	user.Password = "testpassword"

	return user, nil
}

// CreateUser creates a new user
func (m *testDBRepo) CreateUser(user models.User) (string, error) {
	if user.Username == "db_create_user_error" {
		return "", errors.New("error creating user")
	}

	return "1", nil
}
