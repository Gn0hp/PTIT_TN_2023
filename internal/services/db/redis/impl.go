package redis

import (
	"context"
	"encoding/json"
	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
	"time"
)

func (c Client) Ping(ctx context.Context) error {
	_, err := c.redis.Ping(ctx).Result()
	return errors.WithStack(err)
}

func (c Client) Set(ctx context.Context, key string, value interface{}, duration time.Duration) error {
	j, err := json.Marshal(value)
	if err != nil {
		if err == redis.Nil {
			return redis.Nil
		}
		return errors.WithStack(err)
	}
	return errors.WithStack(c.redis.Set(ctx, key, j, duration).Err())
}

// Get attempts to retrieve byte and unmarshal into result
func (c Client) Get(ctx context.Context, key string, value interface{}) error {
	b, err := c.redis.Get(ctx, key).Bytes()
	if err != nil {
		if err == redis.Nil {
			return redis.Nil
		}
		return errors.WithStack(err)
	}
	return errors.WithStack(json.Unmarshal(b, value))
}

func (c Client) HGet(ctx context.Context, hashKey string, values interface{}, emptyJson string, keys ...string) error {
	res := c.redis.HMGet(ctx, hashKey, keys...)
	if res.Err() != nil {
		return res.Err()
	}
	jsonData := "["
	for _, v := range res.Val() {
		tmp, ok := v.(string)
		if !ok {
			tmp = emptyJson
		}
		jsonData += tmp + ","
	}
	jsonData = jsonData[:len(jsonData)-1] + "]"
	b := []byte(jsonData)
	return json.Unmarshal(b, &values)
}

func (c Client) HSet(ctx context.Context, hashKey string, data interface{}) error {
	mapInterface, ok := data.(map[string]interface{})
	if !ok {
		return errors.New("[redis] invalid data: Cannot get map interface")
	}
	return c.redis.HSet(ctx, hashKey, mapInterface).Err()
}

// Exists checks if a key exists in cache
func (c Client) Exists(ctx context.Context, key string) (bool, error) {
	isExist, err := c.redis.Exists(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return false, redis.Nil
		}
		return false, errors.WithStack(err)
	}
	return isExist >= 1, nil
}

func (c Client) Delete(ctx context.Context, key string) error {
	return errors.WithStack(c.redis.Del(ctx, key).Err())
}
