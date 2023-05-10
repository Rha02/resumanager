package driver

import (
	"context"
	"database/sql"
	"time"

	_ "github.com/lib/pq"
	"github.com/redis/go-redis/v9"
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
	db, err := newSQLDatabase(dsn)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(maxOpenDBConns)
	db.SetMaxIdleConns(maxIdleDBConns)
	db.SetConnMaxLifetime(maxDBConnLifetime)

	return &DB{SQL: db}, nil
}

// newDatabase creates a new SQL database connection
func newSQLDatabase(dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		_ = db.Close()
		return nil, err
	}

	return db, nil
}

// KeyValueStore is a wrapper for Redis
type KeyValueStore struct {
	Redis *redis.Client
}

func (m *KeyValueStore) Close() error {
	return m.Redis.Close()
}

const timeout = 3 * time.Second

// ConnectRedis creates a new Redis database connection.
func ConnectRedis(address string, password string) (*KeyValueStore, error) {
	rdb, err := newRedisDatabase(address, password)
	if err != nil {
		return nil, err
	}

	return &KeyValueStore{Redis: rdb}, nil
}

// newDatabase creates a new Redis database connection
func newRedisDatabase(address string, password string) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password,
	})

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	if _, err := client.Ping(ctx).Result(); err != nil {
		_ = client.Close()
		return client, err
	}

	return client, nil
}
