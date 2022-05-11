package main

import (
	"api/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(middleware.CORSMiddleware)
	router(r)
	r.Run(":8090")
}
