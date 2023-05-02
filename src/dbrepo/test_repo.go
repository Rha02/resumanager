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

	user.Email = "user@test.loc"
	user.Username = "testuser"
	user.Password = "testpassword"

	return user, nil
}

// GetUserByUsername gets a user by username
func (m *testDBRepo) GetUserByEmail(email string) (models.User, error) {
	var user models.User
	if email == "db_get_user_error@test.loc" {
		return user, errors.New("error getting user")
	}

	if email == "access_token_error@test.loc" {
		user.Username = "access_token_error"
	} else if email == "refresh_token_error@test.loc" {
		user.Username = "refresh_token_error"
	} else {
		user.Username = "testuser"
	}

	user.ID = 1
	user.Email = email
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
