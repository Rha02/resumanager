package hashservice

type HashRepository interface {
	HashPassword(password string) (string, error)
	ComparePasswords(hashedPassword, password string) bool
}
