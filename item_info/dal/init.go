package dal

import (
	"item_info/dal/ch"
	"item_info/dal/es"
	"item_info/dal/kafka"
	"item_info/dal/pika"
)

func Init() {
	kafka.Init()
	pika.Init()
	es.Init()
	ch.Init()
}

func Close() {
	kafka.Close()
}
