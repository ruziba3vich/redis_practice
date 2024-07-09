package customRedis

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
	"golang.org/x/sync/errgroup"
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

func (r *MyRedis) Get(ctx context.Context, key string) (interface{}, error) {
	return r.redisDb.Get(ctx, key).Result()
}

func (r *MyRedis) Del(ctx context.Context, keys ...string) (response []interface{}, vals int, errs []error) {
	group, c := errgroup.WithContext(ctx)
	for i := range keys {
		group.Go(func() error {
			if resp, err := r.Exists(c, keys[i]); err != nil || resp != 1 {
				errs = append(errs, err)
				return err
			}
			val, err := r.redisDb.Del(c, keys[i]).Result()
			if err != nil {
				errs = append(errs, err)
				return err
			}
			response = append(response, val)
			vals++
			return nil
		})
	}
	if err := group.Wait(); err != nil {
		return nil, 0, errs
	}

	return response, vals, nil
}

func (r *MyRedis) Exists(ctx context.Context, key string) (int64, error) {
	return r.redisDb.Exists(ctx, key).Result()
}

func (r *MyRedis) AddToSet(ctx context.Context, setname string, member string) error {
	return r.redisDb.SAdd(ctx, setname, member).Err()
}

func (r *MyRedis) GetFromSet(ctx context.Context, setname string, member string) (bool, error) {
	return r.redisDb.SIsMember(ctx, setname, member).Result()
}

func (r *MyRedis) RemoveFromSet(ctx context.Context, setname string, member string) error {
	_, err := r.redisDb.SRem(ctx, setname, member).Result()
	return err
}

func (r *MyRedis) AddToHash(ctx context.Context, hashname, key string, val interface{}, duration int) error {
	return r.redisDb.HSet(ctx, hashname, key, val).Err()
}

func (r *MyRedis) RemoveFromHash(ctx context.Context, hashname, key string) error {
	_, err := r.redisDb.HDel(ctx, hashname, key).Result()
	return err
}

func (r *MyRedis) ExistsInHash(ctx context.Context, hashname, key string) (bool, error) {
	return r.redisDb.HExists(ctx, hashname, key).Result()
}

func (r *MyRedis) GetAllFromHash(ctx context.Context, hashname string) (map[string]string, error) {
	return r.redisDb.HGetAll(ctx, hashname).Result()
}

func (r *MyRedis) LeftPush(ctx context.Context, listname, value string) error {
	return r.redisDb.LPush(ctx, listname, value).Err()
}

func (r *MyRedis) RightPush(ctx context.Context, listname, valu string) error {
	return r.redisDb.RPush(ctx, listname, valu).Err()
}

func (r *MyRedis) PopRight(ctx context.Context, listname string) (string, error) {
	return r.redisDb.RPop(ctx, listname).Result()
}

func (r *MyRedis) PopLeft(ctx context.Context, listname string) (string, error) {
	return r.redisDb.LPop(ctx, listname).Result()
}

func (r *MyRedis) ListLength(ctx context.Context, listname string) (int64, error) {
	return r.redisDb.LLen(ctx, listname).Result()
}

func (r *MyRedis) GetRangeElements(ctx context.Context, listname string, from, to int64) ([]string, error) {
	return r.redisDb.LRange(ctx, listname, from, to).Result()
}
