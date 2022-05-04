package dal

import (
	"item_info/dal/kafka"
	"item_info/dal/pika"
)

func Init() {
	kafka.Init()
	pika.Init()
}

func Close() {
	kafka.Close()
}
