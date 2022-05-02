package handlers

import (
	"context"
	"item_info/kitex_gen/base"
	"item_info/kitex_gen/item"
)

func GetItem(ctx context.Context, req *item.GetItemRequest) *item.GetItemResponse {
	resp := item.NewGetItemResponse()
	resp.BaseResp = base.NewBaseResp()

	itemMap := make(map[int64]*item.Item)
	for _, id := range req.GetItemIdList() {
		itemMap[id] = &item.Item{ItemId: id}
	}

	resp.ItemMap = itemMap
	return resp
}
