package cacheservice

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

const timeout = 3 * time.Second

type redisRepo struct {
	rdb *redis.Client
}

func NewRedisRepo(rdb *redis.Client) CacheRepository {
	return &redisRepo{
		rdb: rdb,
	}
}

func (r *redisRepo) Get(key string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	val, _ := r.rdb.Get(ctx, key).Result()
	return val, nil
}

func (r *redisRepo) Set(key string, value string, expiresIn int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	return r.rdb.Set(ctx, key, value, time.Duration(expiresIn)*time.Second).Err()
}

func (r *redisRepo) Close() error {
	return r.rdb.Close()
}
