package ch

import (
	"context"
	"gorm.io/driver/clickhouse"
	"gorm.io/gorm"
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
