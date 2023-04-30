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
	err = initUsersTable(db.SQL)
	if err != nil {
		panic(err)
	}
	log.Println("Users table initialized successfully")

	log.Println("Database initialization completed successfully")
}

// initUsersTable creates the users table in the database
func initUsersTable(db *sql.DB) error {
	stmt := `
		DROP TABLE IF EXISTS users;
		CREATE TABLE users (
			id SERIAL PRIMARY KEY,
			username VARCHAR(50) NOT NULL,
			password_hash VARCHAR(100) NOT NULL,
			created_at TIMESTAMP NOT NULL DEFAULT NOW(),
			updated_at TIMESTAMP NOT NULL DEFAULT NOW()
		);
		CREATE UNIQUE INDEX idx_users_username ON users(username);
	`

	_, err := db.Exec(stmt)
	if err != nil {
		return err
	}

	return nil
}
