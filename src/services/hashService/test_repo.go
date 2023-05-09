package hashservice

import "errors"

type testHashRepo struct{}

func NewTestHashRepo() HashRepository {
	return &testHashRepo{}
}

func (m *testHashRepo) HashPassword(password string) (string, error) {
	if password == "hash_error" {
		return "", errors.New("hash error")
	}

	return password, nil
}

func (m *testHashRepo) ComparePasswords(hashedPassword, password string) bool {
	return password != "wrong_password_error"
}
