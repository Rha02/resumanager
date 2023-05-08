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

// GetUserResumes returns all resumes for a user
func (m *postgresDBRepo) GetUserResumes(userID string) ([]models.Resume, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	var resumes []models.Resume

	stmt := `
		SELECT id, name, file_name, user_id, is_master, size 
		FROM resumes WHERE user_id = $1
	`

	rows, err := m.DB.QueryContext(ctx, stmt, userID)
	if err != nil {
		return resumes, err
	}

	for rows.Next() {
		var resume models.Resume

		if err := rows.Scan(&resume.ID, &resume.Name, &resume.FileName, &resume.UserID, &resume.IsMaster, &resume.Size); err != nil {
			return resumes, err
		}

		resumes = append(resumes, resume)
	}

	return resumes, nil
}

// GetResume returns a resume by ID
func (m *postgresDBRepo) GetResume(id string) (models.Resume, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	var resume models.Resume

	stmt := `
		SELECT id, name, file_name, user_id, is_master, size
		FROM resumes WHERE id = $1
	`

	row := m.DB.QueryRowContext(ctx, stmt, id)
	if err := row.Scan(&resume.ID, &resume.Name, &resume.FileName, &resume.UserID, &resume.IsMaster, &resume.Size); err != nil {
		return resume, err
	}

	return resume, nil
}

// InsertResume inserts a new resume
func (m *postgresDBRepo) InsertResume(resume models.Resume) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	stmt := `
		INSERT INTO resumes (name, file_name, user_id, is_master, size) VALUES ($1, $2, $3, $4, $5)
	`

	_, err := m.DB.ExecContext(ctx, stmt, resume.Name, resume.FileName, resume.UserID, resume.IsMaster, resume.Size)
	if err != nil {
		return err
	}

	return nil
}

// DeleteResume deletes a resume
func (m *postgresDBRepo) DeleteResume(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	stmt := `
		DELETE FROM resumes WHERE id = $1
	`

	_, err := m.DB.ExecContext(ctx, stmt, id)
	if err != nil {
		return err
	}

	return nil
}
