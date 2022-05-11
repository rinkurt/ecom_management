package pika

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
	"item_info/model"
	"testing"
	"time"
)

func TestGet(t *testing.T) {
	Init()
	ctx := context.Background()
	val, err := Get(ctx, "abc")
	fmt.Println(err)
	fmt.Println(errors.Is(err, redis.Nil))
	fmt.Println(val)
}

func TestSetItem(t *testing.T) {
	Init()
	ctx := context.Background()
	err := SetItem(ctx, &model.Item{
		ItemId:       3333,
		UserId:       333333,
		Title:        "hihihiu",
		VideoUrl:     "bb",
		Label:        "sss",
		Status:       0,
		Rate:         0,
		IsEcom:       false,
		ContentLevel: 0,
		CreateTime:   time.Now().Unix(),
		UpdateTime:   time.Now().Unix(),
	})
	fmt.Println(err)
}
