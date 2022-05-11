package ch

import (
	"context"
	"gorm.io/driver/clickhouse"
	"gorm.io/gorm"
	"item_info/kitex_gen/item"
)

var db *gorm.DB

func Init() {
	var err error
	dsn := "tcp://127.0.0.1:9001?debug=true&compress=true&database=default&username=default&password=&read_timeout=10&write_timeout=20"
	db, err = gorm.Open(clickhouse.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

func InsertItemChange(ctx context.Context, ic *ItemChange) error {
	if err := db.WithContext(ctx).Table("item_change").Create(ic).Error; err != nil {
		return err
	}
	return nil
}

func GetItemChangeHistory(ctx context.Context, req *item.GetItemChangeHistoryRequest) ([]*ItemChange, int64, error) {
	q := db.WithContext(ctx).Table("item_change")
	if req.IsSetItemId() {
		q = q.Where("item_id = ?", req.GetItemId())
	}
	if req.IsSetTimeFrom() {
		q = q.Where("time > ?", req.GetTimeFrom())
	}
	if req.IsSetTimeTo() {
		q = q.Where("time < ?", req.GetTimeTo())
	}

	var total int64
	if err := q.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if req.IsSetOrderBy() {
		order := req.GetOrderBy()
		if req.GetOrder() == 0 {
			order += " desc"
		}
		q = q.Order(order)
	}
	if req.IsSetSize() {
		q = q.Limit(int(req.GetSize()))
	}
	if req.IsSetOffset() {
		q = q.Offset(int(req.GetOffset()))
	}

	var res []*ItemChange
	if err := q.Find(&res).Error; err != nil {
		return nil, 0, err
	}

	return res, total, nil
}
