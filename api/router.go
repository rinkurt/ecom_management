package main

import (
	"api/handlers"
	"github.com/gin-gonic/gin"
)

func router(r *gin.Engine) {
	r.POST("/item/get_item_list", handlers.GetItemList)
}
