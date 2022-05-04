package handlers

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/jinzhu/copier"
	"item_info/dal/pika"
	"item_info/kitex_gen/base"
	"item_info/kitex_gen/item"
	"item_info/tools"
)

func GetItem(ctx context.Context, req *item.GetItemRequest) *item.GetItemResponse {
	resp := item.NewGetItemResponse()
	resp.BaseResp = base.NewBaseResp()

	items, err := pika.MGetItem(ctx, req.GetItemIdList())
	if err != nil {
		resp.BaseResp = tools.FillBaseResp(1, err.Error())
		return resp
	}

	itemMap := make(map[int64]*item.Item)
	for id, it := range items {
		itemMap[id] = &item.Item{}
		if err := copier.Copy(itemMap[id], it); err != nil {
			klog.CtxErrorf(ctx, "copier failed: %v", err)
		}
	}

	resp.ItemMap = itemMap
	return resp
}
