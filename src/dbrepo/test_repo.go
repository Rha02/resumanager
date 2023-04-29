package dbrepo

import "github.com/Rha02/resumanager/src/models"

type testDBRepo struct{}

// NewTestDBRepo creates a new test database repository
func NewTestDBRepo() DatabaseRepository {
	return &testDBRepo{}
}

func (m *testDBRepo) GetUserByID(id string) (models.User, error) {
	return models.User{}, nil
}

func (m *testDBRepo) GetUserByUsername(username string) (models.User, error) {
	return models.User{}, nil
}

func (m *testDBRepo) CreateUser(user models.User) (int, error) {
	return 1, nil
}
