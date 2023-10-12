package repositories

import (
	"PTIT_TN/internal/services/db/redis"
	"context"
	"time"
)

type CacheRepo interface {
	Ping(ctx context.Context) error
	Get(ctx context.Context, key string, value interface{}) error
	Set(ctx context.Context, key string, value interface{}, duration time.Duration) error
	HGet(ctx context.Context, hashKey string, values interface{}, emptyJson string, keys ...string) error
	HSet(ctx context.Context, hashKey string, data interface{}) error
	Exists(ctx context.Context, key string) (bool, error)
	Delete(ctx context.Context, key string) error
}

func NewCacheRepo(client *redis.Client) CacheRepo {
	return client
}
