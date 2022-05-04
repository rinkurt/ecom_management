package tools

import (
	"encoding/json"
	"item_info/kitex_gen/base"
)

func FillBaseResp(code int64, msg string) *base.BaseResp {
	return &base.BaseResp{StatusCode: code, StatusMessage: msg}
}

func ConvertMapSlice(in []map[string]interface{}, out interface{}) error {
	bytes, err := json.Marshal(in)
	if err != nil {
		return err
	}
	err = json.Unmarshal(bytes, out)
	if err != nil {
		return err
	}
	return nil
}
