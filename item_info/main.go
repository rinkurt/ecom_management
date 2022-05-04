package main

import (
	"item_info/dal"
	item "item_info/kitex_gen/item/itemservice"
	"log"
)

func main() {
	defer dal.Close()
	dal.Init()

	svr := item.NewServer(new(ItemServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
