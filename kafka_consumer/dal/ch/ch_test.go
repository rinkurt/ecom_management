package ch

import (
	"context"
	_ "github.com/ClickHouse/clickhouse-go"
	"testing"
)

func TestInsertItemChange(t *testing.T) {
	Init()
	InsertItemChange(context.Background(), &ItemChange{
		ItemId:     112,
		Time:       0,
		ItemBefore: "{}",
		ItemAfter:  "{\"title\":\"mmmm\"}",
	})
}
