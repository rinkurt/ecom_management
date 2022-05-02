package main

import (
	"item_info/dal/pika"
	item "item_info/kitex_gen/item/itemservice"
	"log"
)

func main() {
	pika.Init()

	svr := item.NewServer(new(ItemServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
