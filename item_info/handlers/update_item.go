package handlers

import (
	"context"
	"item_info/dal/pika"
	"item_info/kitex_gen/base"
	"item_info/kitex_gen/item"
	"item_info/tools"
)

func UpdateItem(ctx context.Context, req *item.UpdateItemRequest) *item.UpdateItemResponse {
	resp := item.NewUpdateItemResponse()
	resp.BaseResp = base.NewBaseResp()

	it, err := pika.GetItem(ctx, req.GetItemId())
	if err != nil {
		resp.BaseResp = tools.FillBaseResp(1, err.Error())
		return resp
	}

	if req.IsSetUserId() {
		it.UserId = req.GetUserId()
	}

	err = pika.SetItem(ctx, it)
	if err != nil {
		resp.BaseResp = tools.FillBaseResp(1, err.Error())
		return resp
	}

	return resp
}
