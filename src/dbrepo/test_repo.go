package dbrepo

import (
	"errors"
	"strconv"

	"github.com/Rha02/resumanager/src/models"
)

type testDBRepo struct {
	users []models.User
}

// NewTestDBRepo creates a new test database repository
func NewTestDBRepo() DatabaseRepository {
	return &testDBRepo{
		users: []models.User{},
	}
}

// GetUserByID gets a user by ID
func (m *testDBRepo) GetUserByID(id string) (models.User, error) {
	var user models.User

	userID, err := strconv.Atoi(id)
	if err != nil {
		return user, err
	}

	for _, u := range m.users {
		if u.ID == userID {
			user = u
			break
		}
	}

	return user, nil
}

// GetUserByUsername gets a user by username
func (m *testDBRepo) GetUserByUsername(username string) (models.User, error) {
	for _, u := range m.users {
		if u.Username == username {
			return u, nil
		}
	}

	return models.User{}, errors.New("user not found")
}

// CreateUser creates a new user
func (m *testDBRepo) CreateUser(user models.User) (string, error) {
	user.ID = len(m.users) + 1
	m.users = append(m.users, user)

	if user.Username == "error" {
		return "", errors.New("error creating user")
	}

	return strconv.Itoa(user.ID), nil
}
