package pika

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
	"testing"
)

func TestGet(t *testing.T) {
	Init()
	ctx := context.Background()
	val, err := Get(ctx, "abc")
	fmt.Println(err)
	fmt.Println(errors.Is(err, redis.Nil))
	fmt.Println(val)
}
