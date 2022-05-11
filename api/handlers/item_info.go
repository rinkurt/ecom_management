package handlers

import (
	"api/kitex_gen/item"
	"api/kitex_gen/item/itemservice"
	"api/tools"
	"context"
	"github.com/cloudwego/kitex/client"
	"github.com/gin-gonic/gin"
	"log"
)

var itemCli itemservice.Client

func init() {
	cli, err := itemservice.NewClient("hello", client.WithHostPorts("0.0.0.0:8888"))
	if err != nil {
		log.Fatal(err)
	}
	itemCli = cli
}

func GetItemList(c *gin.Context) {
	c.Set("Handler", "GetItemList")
	ctx := context.Background()
	req := item.NewGetItemListRequest()

	err := c.ShouldBindJSON(&req)
	if err != nil {
		tools.HandleError(c, tools.StatusBindError, err)
		return
	}

	resp, err := itemCli.GetItemList(ctx, req)
	if err != nil {
		tools.HandleError(c, tools.StatusRequestError, err)
		return
	}

	c.JSON(200, tools.NewDataMap(resp))
}
