package tools

import "item_info/kitex_gen/base"

func FillBaseResp(code int64, msg string) *base.BaseResp {
	return &base.BaseResp{StatusCode: code, StatusMessage: msg}
}
