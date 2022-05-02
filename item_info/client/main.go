package main

import (
	"context"
	"github.com/cloudwego/kitex/client"
	"item_info/kitex_gen/item"
	"item_info/kitex_gen/item/itemservice"
	"log"
	"time"
)

func main() {
	client, err := itemservice.NewClient("hello", client.WithHostPorts("0.0.0.0:8888"))
	if err != nil {
		log.Fatal(err)
	}
	for {
		req := &item.GetItemRequest{ItemIdList: []int64{1234}}
		resp, err := client.GetItem(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(resp)
		time.Sleep(time.Second)
	}
}
