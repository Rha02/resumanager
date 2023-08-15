package hashservice

import "golang.org/x/crypto/bcrypt"

const (
	// Cost is the cost of the bcrypt hashing algorithm
	cost = bcrypt.DefaultCost
)

type bcryptRepo struct{}

func NewBcryptRepo() HashRepository {
	return &bcryptRepo{}
}

// HashPassword hashes a password
func (m *bcryptRepo) HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

// ComparePasswords compares a hashed password with a plain text password
func (m *bcryptRepo) ComparePasswords(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
