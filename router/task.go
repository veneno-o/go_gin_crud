package router

import (
	"github.com/gin-gonic/gin"
	"go_gin/service"
	"net/http"
	"os"
)

func RouterHandler(ginServe *gin.Engine) {
	root, _ := os.Getwd()

	ginServe.Static("/css", root+"/static/css")
	ginServe.Static("/fonts", root+"/static/fonts")
	ginServe.Static("/js", root+"/static/js")

	ginServe.LoadHTMLGlob("templates/*")

	ginServe.GET("/", service.GetPageHandler)
	routerGroup := ginServe.Group("/api")
	{
		//查询
		routerGroup.GET("/", service.GetAllTaskHandler)

		//添加
		routerGroup.POST("/", service.AddTaskHandler)

		//修改
		routerGroup.PUT("/:id", service.UpdataTaskHandler)

		//删除
		routerGroup.DELETE("/:id", service.DelTaskHandler)
	}

	// no match response 404.html
	ginServe.NoRoute(func(ctx *gin.Context) {
		ctx.HTML(http.StatusNotFound, "404.html", gin.H{})
	})
}
