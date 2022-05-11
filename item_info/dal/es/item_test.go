package es

import (
	"context"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"github.com/sanity-io/litter"
	"item_info/kitex_gen/item"
	"item_info/model"
	"testing"
	"time"
)

func TestUpsertItem(t *testing.T) {
	Init()
	it := &model.Item{
		ItemId:       2222,
		UserId:       12345,
		Title:        "ddddd",
		VideoUrl:     "sss",
		Label:        "ttt",
		Status:       1,
		Rate:         2,
		IsEcom:       false,
		ContentLevel: 0,
		CreateTime:   time.Now().Unix(),
		UpdateTime:   time.Now().Unix(),
	}
	err := UpsertItem(context.Background(), it)
	fmt.Println(err)
}

func TestSearchItem(t *testing.T) {
	Init()
	res, total, err := SearchItem(context.Background(), &item.GetItemListRequest{
		ItemId:         nil,
		Title:          nil,
		CreateTimeFrom: nil,
		CreateTimeTo:   nil,
		Label:          nil,
		UserId:         nil,
		Status:         nil,
		Rate:           nil,
		ContentLevel:   nil,
		Order:          1,
		OrderBy:        "item_id",
		Size:           20,
		Offset:         thrift.Int64Ptr(0),
	})
	fmt.Println(err)
	fmt.Println(total)
	fmt.Println(litter.Sdump(res))
}
