package handlers

import (
	"context"
	"github.com/jinzhu/copier"
	"item_info/dal/kafka"
	"item_info/dal/pika"
	"item_info/kitex_gen/base"
	"item_info/kitex_gen/item"
	"item_info/model"
	"item_info/tools"
	"time"
)

func CreateItem(ctx context.Context, req *item.CreateItemRequest) *item.CreateItemResponse {
	resp := item.NewCreateItemResponse()
	resp.BaseResp = base.NewBaseResp()

	it := &model.Item{}

	err := copier.CopyWithOption(it, req, copier.Option{IgnoreEmpty: true})
	if err != nil {
		resp.BaseResp = tools.FillBaseResp(1, err.Error())
		return resp
	}

	err = pika.SetItem(ctx, it)
	if err != nil {
		resp.BaseResp = tools.FillBaseResp(1, err.Error())
		return resp
	}

	kafka.SendItemChange(ctx, &model.ItemChange{
		ItemId:     it.ItemId,
		Time:       time.Now().Unix(),
		ItemBefore: nil,
		ItemAfter:  it,
	})

	return resp
}
