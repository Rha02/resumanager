package cacheservice

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
	return r.cache[key], nil
}

// Set(key, value) sets the value of the key
func (r *testCacheRepo) Set(key string, value string) error {
	r.cache[key] = value
	return nil
}
