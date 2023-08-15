package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/Rha02/resumanager/src/driver"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	connStr := os.Getenv("DB_CONNECTION")

	log.Println("Connecting to database...")
	db, err := driver.ConnectSQL(connStr)
	if err != nil {
		panic(err)
	}
	log.Println("Connected to database successfully")

	defer db.Close()

	log.Println("Initializing users table...")
	if err = initUsersTable(db.SQL); err != nil {
		panic(err)
	}
	log.Println("Users table initialized successfully")

	log.Println("Initializing resumes table...")
	if err = initResumesTable(db.SQL); err != nil {
		panic(err)
	}
	log.Println("Resumes table initialized successfully")

	log.Println("Database initialization completed successfully")
}

// initUsersTable creates the users table in the database
func initUsersTable(db *sql.DB) error {
	stmt := `
		DROP TABLE IF EXISTS users;
		CREATE TABLE users (
			id SERIAL PRIMARY KEY,
			email VARCHAR(254) NOT NULL,
			username VARCHAR(50) NOT NULL,
			password_hash VARCHAR(100) NOT NULL,
			created_at TIMESTAMP NOT NULL DEFAULT NOW(),
			updated_at TIMESTAMP NOT NULL DEFAULT NOW()
		);
		CREATE UNIQUE INDEX users_email_idx ON users (email);
	`

	_, err := db.Exec(stmt)
	if err != nil {
		return err
	}

	return nil
}

// initResumesTable creates the resumes table in the database
func initResumesTable(db *sql.DB) error {
	stmt := `
		DROP TABLE IF EXISTS resumes;
		CREATE TABLE resumes (
			id SERIAL PRIMARY KEY,
			user_id INT NOT NULL,
			name VARCHAR(50) NOT NULL,
			file_name VARCHAR(100) NOT NULL,
			is_master BOOLEAN NOT NULL DEFAULT FALSE,
			size INT NOT NULL,
			FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
		);
	`

	_, err := db.Exec(stmt)
	if err != nil {
		return err
	}
	return nil
}
