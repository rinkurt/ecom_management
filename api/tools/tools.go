package tools

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"log"
)

const (
	StatusBindError    = 1
	StatusRequestError = 2
)

func NewErrorMap(code int, err error) gin.H {
	if err == nil {
		return NewDataMap(nil)
	}
	return gin.H{
		"code": code,
		"msg":  err.Error(),
	}
}

func NewDataMap(data interface{}) gin.H {
	return gin.H{
		"code": 0,
		"msg":  "success",
		"data": data,
	}
}

func HandleError(c *gin.Context, code int, err error) {
	if handler, ok := c.Get("Handler"); ok {
		err = errors.WithMessage(err, fmt.Sprintf("%v Handler", handler))
	}
	log.Printf("Error: %v", err)
	c.JSON(200, NewErrorMap(code, err))
}
