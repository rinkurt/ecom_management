package ch

import (
	"context"
	"fmt"
	_ "github.com/ClickHouse/clickhouse-go"
	"github.com/apache/thrift/lib/go/thrift"
	"github.com/sanity-io/litter"
	"item_info/kitex_gen/item"
	"testing"
)

func TestGetItemChangeHistory(t *testing.T) {
	Init()
	history, total, err := GetItemChangeHistory(context.Background(), &item.GetItemChangeHistoryRequest{
		ItemId:   nil,
		TimeFrom: nil,
		TimeTo:   nil,
		Order:    nil,
		OrderBy:  thrift.StringPtr("item_id"),
		Size:     nil,
		Offset:   nil,
	})
	fmt.Println(litter.Sdump(history), total, err)
}

func TestInsertItemChange(t *testing.T) {
	Init()
	InsertItemChange(context.Background(), &ItemChange{
		ItemId:     112,
		Time:       0,
		ItemBefore: "{}",
		ItemAfter:  "{\"title\":\"mmmm\"}",
	})
}
