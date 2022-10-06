package handlers

import (
	"context"
	"item_info/constdef"
	"item_info/dal/pika"
	"item_info/kitex_gen/base"
	"item_info/kitex_gen/item"
	"item_info/tools"
)

func IncrCount(ctx context.Context, req *item.IncrCountRequest) *item.IncrCountResponse {
	resp := item.NewIncrCountResponse()
	resp.BaseResp = base.NewBaseResp()

	countType := ""
	switch req.GetCountType() {
	case item.CountType_Play:
		countType = constdef.CountTypePlay
	case item.CountType_Like:
		countType = constdef.CountTypeLike
	case item.CountType_Share:
		countType = constdef.CountTypeShare
	case item.CountType_Comment:
		countType = constdef.CountTypeComment
	default:
		resp.BaseResp = tools.FillBaseResp(2, "unknown count type")
		return resp
	}

	err := pika.IncrCount(ctx, req.GetItemId(), countType, req.GetIncr())
	if err != nil {
		resp.BaseResp = tools.FillBaseResp(1, err.Error())
		return resp
	}

	return resp
}
