package cacheservice

import "errors"

type testCacheRepo struct {
	cache map[string]string
}

func NewTestCacheRepo() CacheRepository {
	return &testCacheRepo{
		cache: make(map[string]string),
	}
}

// Get(key) returns the value of the key
func (r *testCacheRepo) Get(key string) (string, error) {
	if key == "cache_error" {
		return "", errors.New("error")
	}
	return r.cache[key], nil
}

// Set(key, value) sets the value of the key
func (r *testCacheRepo) Set(key string, value string) error {
	if key == "cache_error" {
		return errors.New("error")
	}
	r.cache[key] = value
	return nil
}
