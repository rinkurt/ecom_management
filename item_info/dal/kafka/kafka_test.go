package kafka

import (
	"context"
	"item_info/model"
	"testing"
	"time"
)

func TestProduce(t *testing.T) {
	defer Close()
	Init()

	Produce(context.Background(), "hello2")
}

func TestSendItemChange(t *testing.T) {
	defer Close()
	Init()

	SendItemChange(context.Background(), &model.ItemChange{
		ItemId:     123,
		Time:       time.Now().Unix(),
		ItemBefore: nil,
		ItemAfter: &model.Item{
			ItemId:       123,
			UserId:       0,
			Title:        "ggggg",
			VideoUrl:     "heheheh",
			Label:        "",
			Status:       0,
			Rate:         0,
			IsEcom:       false,
			ContentLevel: 0,
			CreateTime:   0,
			UpdateTime:   0,
		},
	})
}
