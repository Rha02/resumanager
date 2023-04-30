package driver

import (
	"database/sql"
	"time"

	_ "github.com/lib/pq"
)

type DB struct {
	SQL *sql.DB
}

// Close closes the database connection
func (db *DB) Close() error {
	return db.SQL.Close()
}

const (
	maxOpenDBConns    = 1
	maxIdleDBConns    = 1
	maxDBConnLifetime = 1 * time.Minute
)

// ConnectSQL creates a new database connection.
// It returns a pointer to the DB struct and an error.
// If the error is not nil, the db pointer will be nil.
func ConnectSQL(dsn string) (*DB, error) {
	db, err := newDatabase(dsn)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(maxOpenDBConns)
	db.SetMaxIdleConns(maxIdleDBConns)
	db.SetConnMaxLifetime(maxDBConnLifetime)

	return &DB{SQL: db}, nil
}

// newDatabase creates a new database connection
func newDatabase(dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	return db, nil
}
