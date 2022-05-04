package es

import (
	"context"
	"github.com/spf13/cast"
	"kafka_consumer/model"
)

const IndexItem = "item"

func UpsertItem(ctx context.Context, it *model.Item) error {
	return Upsert(ctx, IndexItem, cast.ToString(it.ItemId), it)
}
