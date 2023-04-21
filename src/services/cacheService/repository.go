package cacheservice

type CacheRepository interface {
	Get(key string) (string, error)
	Set(key string, value string) error
}
