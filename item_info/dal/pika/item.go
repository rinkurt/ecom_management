package pika

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/pkg/errors"
	"item_info/model"
)

func KeyItem(itemId int64) string {
	return fmt.Sprintf("item_%d", itemId)
}

func MGetItem(ctx context.Context, itemIdList []int64) (map[int64]*model.Item, error) {
	keys := make([]string, 0)
	for _, id := range itemIdList {
		keys = append(keys, KeyItem(id))
	}
	vals, err := rdb.MGet(ctx, keys...).Result()
	if err != nil {
		return nil, errors.WithMessage(err, "MGetItem: pika mget failed")
	}
	res := make(map[int64]*model.Item)
	for _, val := range vals {
		// ignore nil
		if valStr, ok := val.(string); ok {
			it := &model.Item{}
			err := json.Unmarshal([]byte(valStr), &it)
			if err != nil {
				klog.CtxWarnf(ctx, "Unmarshal item failed: err=%v", err)
				continue
			}
			res[it.ItemId] = it
		}
	}
	return res, nil
}

func GetItem(ctx context.Context, itemId int64) (*model.Item, error) {
	itemMap, err := MGetItem(ctx, []int64{itemId})
	if err != nil {
		return nil, err
	}
	item, ok := itemMap[itemId]
	if !ok {
		return nil, errors.New("item not found")
	}
	return item, nil
}

func SetItem(ctx context.Context, item *model.Item) error {
	bytes, err := json.Marshal(item)
	if err != nil {
		return errors.WithMessage(err, "SetItem, marshal failed")
	}
	err = rdb.Set(ctx, KeyItem(item.ItemId), string(bytes), 0).Err()
	if err != nil {
		return errors.WithMessage(err, "SetItem, pika set failed")
	}
	return nil
}
