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

func NewRedisRepo(address string, password string) CacheRepository {
	client := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password,
	})

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	_, err := client.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}

	return &redisRepo{
		rdb: client,
	}
}

func (r *redisRepo) Get(key string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	val, err := r.rdb.Get(ctx, key).Result()
	if err != nil {
		return "", nil
	}

	return val, nil
}

func (r *redisRepo) Set(key string, value string, expiresIn int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	err := r.rdb.Set(ctx, key, value, time.Duration(expiresIn)*time.Second).Err()
	if err != nil {
		return err
	}

	return nil
}

func (r *redisRepo) Close() error {
	return r.rdb.Close()
}
