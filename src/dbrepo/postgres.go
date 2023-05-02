package dbrepo

import (
	"context"
	"database/sql"
	"time"

	"github.com/Rha02/resumanager/src/models"
)

const timeout = 3 * time.Second

type postgresDBRepo struct {
	DB *sql.DB
}

func NewPostgresRepo(db *sql.DB) DatabaseRepository {
	return &postgresDBRepo{
		DB: db,
	}
}

// GetUserByID returns a user by ID
func (m *postgresDBRepo) GetUserByID(id string) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	var user models.User

	stmt := `SELECT id, email, username, password_hash FROM users WHERE id = $1`

	row := m.DB.QueryRowContext(ctx, stmt, id)

	if err := row.Scan(&user.ID, &user.Email, &user.Username, &user.Password); err != nil {
		return user, err
	}

	return user, nil
}

// GetUserByUsername returns a user by username
func (m *postgresDBRepo) GetUserByEmail(email string) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	var user models.User

	stmt := `
		SELECT id, email, username, password_hash
		FROM users
		WHERE email = $1
	`

	row := m.DB.QueryRowContext(ctx, stmt, email)

	if err := row.Scan(&user.ID, &user.Email, &user.Username, &user.Password); err != nil {
		return user, err
	}

	return user, nil
}

// CreateUser creates a new user
func (m *postgresDBRepo) CreateUser(user models.User) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	stmt := `
		INSERT INTO users (email, username, password_hash) VALUES ($1, $2, $3) RETURNING id
	`

	var newID string

	if err := m.DB.QueryRowContext(ctx, stmt, user.Email, user.Username, user.Password).Scan(&newID); err != nil {
		return "", err
	}

	return newID, nil
}
