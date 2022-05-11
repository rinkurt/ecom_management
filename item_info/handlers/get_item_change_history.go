package handlers

import (
	"context"
	"github.com/jinzhu/copier"
	"item_info/dal/ch"
	"item_info/kitex_gen/base"
	"item_info/kitex_gen/item"
	"item_info/tools"
)

func GetItemChangeHistory(ctx context.Context, req *item.GetItemChangeHistoryRequest) *item.GetItemChangeHistoryResponse {
	resp := item.NewGetItemChangeHistoryResponse()
	resp.BaseResp = base.NewBaseResp()

	res, total, err := ch.GetItemChangeHistory(ctx, req)
	if err != nil {
		resp.BaseResp = tools.FillBaseResp(1, err.Error())
		return resp
	}

	for _, ic := range res {
		resIc := &item.ItemChange{}
		copier.Copy(resIc, ic)
		resp.History = append(resp.History, resIc)
	}

	resp.Total = total
	return resp
}
