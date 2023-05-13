package storage

import (
	"context"

	"github.com/redis/go-redis/v9"
)

func NewRedisClient(addr string) *redis.Client {
	rdb := redis.NewClient(
		&redis.Options{
			Addr:     addr,
			Password: "", // no password set
			DB:       0,  // use default DB
		},
	)
	return rdb
}

func (r *Repository) IncreaseCache(ctx context.Context, key string, val int64) (int64, error) {
	res, err := r.cache.IncrBy(ctx, key, val).Result()
	if err != nil {
		return 0, err
	}

	return res, nil
}
