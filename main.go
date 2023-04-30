package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	e := gin.Default()
	e.GET("/", func(ctx *gin.Context) {
		ctx.JSONP(http.StatusOK, "Hello Gin")
	})
	e.Run(":8080")
}
