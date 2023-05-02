package cacheservice

import "errors"

type testRepo struct {
	cache map[string]string
}

func NewTestRepo() CacheRepository {
	return &testRepo{
		cache: make(map[string]string),
	}
}

// Get(key) returns the value of the key
func (r *testRepo) Get(key string) (string, error) {
	if key == "cache_error" {
		return "", errors.New("error")
	}
	return r.cache[key], nil
}

// Set(key, value) sets the value of the key
func (r *testRepo) Set(key string, value string, expiresIn int64) error {
	if key == "cache_error" {
		return errors.New("error")
	}
	r.cache[key] = value
	return nil
}

// Close() closes the connection to the cache
func (r *testRepo) Close() error {
	return nil
}
