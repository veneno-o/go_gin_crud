package main

import (
	"github.com/gin-gonic/gin"
	"go_gin/dao"
	"gorm.io/gorm"
	"log"
	"net/http"
)

var DB *gorm.DB

func init() {
	DB = dao.Connect()
}

func main() {
	ginServe := gin.Default()

	routerGroup := ginServe.Group("/")
	{
		//获取一个任务OK
		routerGroup.GET("/:id", func(ctx *gin.Context) {
			id := ctx.Param("id")
			task := dao.Task{}
			DB.Debug().First(&task, id)
			ctx.JSON(http.StatusOK, gin.H{
				"code": 2000,
				"data": gin.H{
					"content": task.Content,
					"status":  task.Status,
				},
			})
		})

		//提交一个任务
		routerGroup.POST("/:content", func(ctx *gin.Context) {
			content := ctx.Param("content")
			task := dao.Task{Content: &content}
			DB.Debug().Create(&task)
			ctx.JSON(http.StatusOK, gin.H{
				"code": 2000,
				"msg":  "",
			})
		})

		//修改一个任务(修改任务状态)
		routerGroup.PUT("/:id", func(ctx *gin.Context) {
			id := ctx.Param("id")
			//id_, _ := strconv.Atoi(id)
			//task := dao.Task{
			//	Model: gorm.Model{ID: uint(id_)},
			//}
			//tx := DB.Debug().Select(gin.H{
			//	"id": id,
			//})
			log.Printf("id:%t", id)
		})

		//删除一个任务
		routerGroup.DELETE("/:id", func(context *gin.Context) {
			log.Println("delete")
		})
	}
	ginServe.NoRoute(func(ctx *gin.Context) {
		ctx.HTML(http.StatusNotFound, "404.html", gin.H{})
	})
	ginServe.Run(":8080")
}
