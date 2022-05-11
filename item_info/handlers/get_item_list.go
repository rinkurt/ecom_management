package handlers

import (
	"context"
	"github.com/jinzhu/copier"
	"item_info/dal/es"
	"item_info/kitex_gen/base"
	"item_info/kitex_gen/item"
	"item_info/tools"
)

func GetItemList(ctx context.Context, req *item.GetItemListRequest) *item.GetItemListResponse {
	resp := item.NewGetItemListResponse()
	resp.BaseResp = base.NewBaseResp()

	items, total, err := es.SearchItem(ctx, req)
	if err != nil {
		resp.BaseResp = tools.FillBaseResp(1, err.Error())
		return resp
	}

	for _, it := range items {
		resIt := &item.Item{}
		copier.Copy(resIt, it)
		resp.ItemList = append(resp.ItemList, resIt)
	}

	resp.Total = total
	return resp
}
