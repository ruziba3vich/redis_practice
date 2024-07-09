package repo

import "context"

type RedisClient interface {
	Put(ctx context.Context, key string, value interface{}, expiration int) error
	Get(ctx context.Context, key string) (interface{}, error)
	Del(ctx context.Context, keys ...string) (response []interface{}, vals int, errs []error)
	Exists(ctx context.Context, key string) (int64, error)
	AddToSet(ctx context.Context, setname string, member string) error
	GetFromSet(ctx context.Context, setname string, member string) (bool, error)
	RemoveFromSet(ctx context.Context, setname string, member string) error
	AddToHash(ctx context.Context, hashname, key string, val interface{}, duration int) error
	RemoveFromHash(ctx context.Context, hashname, key string) error
	ExistsInHash(ctx context.Context, hashname, key string) (bool, error)
	GetAllFromHash(ctx context.Context, hashname string) (map[string]string, error)
	LeftPush(ctx context.Context, listname, value string) error
	RightPush(ctx context.Context, listname, value string) error
	PopRight(ctx context.Context, listname string) (string, error)
	PopLeft(ctx context.Context, listname string) (string, error)
	ListLength(ctx context.Context, listname string) (int64, error)
	GetRangeElements(ctx context.Context, listname string, from, to int64) ([]string, error)
}
