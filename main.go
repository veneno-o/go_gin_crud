package main

import (
	"github.com/gin-gonic/gin"
	"go_gin/dao"
	"go_gin/router"
)

func init() {
	dao.Connect()
}

func main() {
	ginServe := gin.Default()
	router.RouterHandler(ginServe)
	ginServe.Run(":8000")
}
