package main

import (
	"context"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"github.com/cloudwego/kitex/client"
	"github.com/sanity-io/litter"
	"item_info/kitex_gen/item"
	"item_info/kitex_gen/item/itemservice"
	"log"
	"testing"
)

var cli itemservice.Client
var ctx = context.Background()

func init() {
	var err error
	cli, err = itemservice.NewClient("hello", client.WithHostPorts("0.0.0.0:8888"))
	if err != nil {
		log.Fatal(err)
	}
}

func TestCreateItem(t *testing.T) {
	for i := 10000; i < 10050; i++ {
		resp, err := cli.CreateItem(ctx, &item.CreateItemRequest{
			ItemId:   int64(i),
			UserId:   int64(i + 10000),
			Title:    fmt.Sprintf("id%v", i),
			VideoUrl: fmt.Sprintf("url_%v", i),
			Label:    "sample_label",
		})
		fmt.Println(err)
		fmt.Println(litter.Sdump(resp))
	}
}

func TestGetItem(t *testing.T) {
	resp, err := cli.GetItem(ctx, &item.GetItemRequest{ItemIdList: []int64{123}})
	fmt.Println(err)
	fmt.Println(litter.Sdump(resp))
}

func TestUpdateItem(t *testing.T) {
	resp, err := cli.UpdateItem(ctx, &item.UpdateItemRequest{
		ItemId: 10001,
		IsEcom: thrift.BoolPtr(true),
	})
	fmt.Println(err)
	fmt.Println(litter.Sdump(resp))
}

func TestGetItemChange(t *testing.T) {
	resp, err := cli.GetItemChangeHistory(ctx, &item.GetItemChangeHistoryRequest{})
	fmt.Println(err)
	fmt.Println(litter.Sdump(resp))
}
