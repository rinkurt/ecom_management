package pika

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
)

var rdb *redis.Client

func Init() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:9221",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

func Get(ctx context.Context, key string) (string, error) {
	val, err := rdb.Get(ctx, key).Result()
	if err != nil {
		return "", errors.WithMessagef(err, "pika get failed, key=%v", key)
	}
	return val, nil
}

func Set(ctx context.Context, key, val string) error {
	err := rdb.Set(ctx, key, val, 0).Err()
	if err != nil {
		return errors.WithMessagef(err, "pika set failed, key=%v, val=%v", key, val)
	}
	return nil
}

func IsNil(err error) bool {
	return errors.Is(err, redis.Nil)
}
