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
	resp, err := cli.CreateItem(ctx, &item.CreateItemRequest{
		ItemId:   123,
		UserId:   456,
		Title:    "hahaha",
		VideoUrl: "sample_video_url",
		Label:    "sample_label",
	})
	fmt.Println(err)
	fmt.Println(litter.Sdump(resp))
}

func TestGetItem(t *testing.T) {
	resp, err := cli.GetItem(ctx, &item.GetItemRequest{ItemIdList: []int64{123}})
	fmt.Println(err)
	fmt.Println(litter.Sdump(resp))
}

func TestUpdateItem(t *testing.T) {
	resp, err := cli.UpdateItem(ctx, &item.UpdateItemRequest{
		ItemId: 123,
		Title:  thrift.StringPtr("hahaha2"),
		Label:  thrift.StringPtr("sample_label2"),
	})
	fmt.Println(err)
	fmt.Println(litter.Sdump(resp))
}
