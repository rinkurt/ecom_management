package model

type ItemChange struct {
	ItemId     int64 `json:"item_id"`
	Time       int64 `json:"time"`
	ItemBefore *Item `json:"item_before"`
	ItemAfter  *Item `json:"item_after"`
}

type Item struct {
	ItemId       int64  `thrift:"item_id,1" json:"item_id"`
	UserId       int64  `thrift:"user_id,2" json:"user_id"`
	Title        string `thrift:"title,3" json:"title"`
	VideoUrl     string `thrift:"video_url,4" json:"video_url"`
	Label        string `thrift:"label,5" json:"label"`
	Status       int64  `thrift:"status,6" json:"status"`
	Rate         int64  `thrift:"rate,7" json:"rate"`
	IsEcom       bool   `thrift:"is_ecom,8" json:"is_ecom"`
	ContentLevel int64  `thrift:"content_level,9" json:"content_level"`
	CreateTime   int64  `thrift:"create_time,11" json:"create_time"`
	UpdateTime   int64  `thrift:"update_time,12" json:"update_time"`
}
