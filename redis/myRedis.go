package customRedis

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

type MyRedis struct {
	redisDb *redis.Client
}

func NewMyRedis(redisDb *redis.Client) *MyRedis {
	return &MyRedis{
		redisDb: redisDb,
	}
}

func (r *MyRedis) Put(ctx context.Context, key string, value interface{}, expiration int) error {
	return r.redisDb.Set(ctx, key, value, time.Duration(expiration)).Err()
}

func (r *MyRedis) Get
