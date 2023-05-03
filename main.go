package main

import (
	"github.com/gin-gonic/gin"
	"go_gin/dao"
	"gorm.io/gorm"
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
		routerGroup.GET("/", func(ctx *gin.Context) {
			var task []dao.Task

			tx := DB.Debug().Find(&task)
			if tx.RowsAffected == 0 {
				ctx.JSON(http.StatusOK, gin.H{
					"code": 2001,
					"msg":  "没有记录",
					"data": nil,
				})
				return
			}
			ctx.JSON(http.StatusOK, gin.H{
				"code": 2000,
				"msg":  "OK",
				"data": task,
			})
		}) //获取一个任务OK
		routerGroup.GET("/:id", func(ctx *gin.Context) {
			var task dao.Task
			id := ctx.Param("id")
			tx := DB.Debug().First(&task, id)
			if tx.RowsAffected == 0 {
				ctx.JSON(http.StatusOK, gin.H{
					"code": 2001,
					"msg":  "记录不存在",
					"data": nil,
				})
				return
			}
			ctx.JSON(http.StatusOK, gin.H{
				"code": 2000,
				"msg":  "OK",
				"data": gin.H{
					"id":      task.ID,
					"content": task.Content,
					"status":  task.Status,
				},
			})
		})

		//提交一个任务OK
		routerGroup.POST("/:content", func(ctx *gin.Context) {
			content := ctx.Param("content")
			task := dao.Task{Content: &content}
			DB.Debug().Create(&task)
			ctx.JSON(http.StatusOK, gin.H{
				"code": 2000,
				"msg":  "新增成功",
				"data": nil,
			})
		})

		//修改一个任务(修改任务状态)
		routerGroup.PUT("/:id/:state", func(ctx *gin.Context) {
			var task dao.Task
			state := ctx.Param("state")
			id := ctx.Param("id")
			//1.查询
			tx := DB.Debug().First(&task, id)
			if tx.RowsAffected == 0 {
				ctx.JSON(http.StatusOK, gin.H{
					"code": 2001,
					"msg":  "记录不存在",
					"data": nil,
				})
				return
			}
			//2.修改
			DB.Debug().Model(&task).Update("Status", state)
			//3.返回json
			ctx.JSON(http.StatusOK, gin.H{
				"code": 2000,
				"msg":  "修改成功",
				"data": nil,
			})
		})

		//删除一个任务
		routerGroup.DELETE("/:id", func(ctx *gin.Context) {
			var task dao.Task
			id := ctx.Param("id")
			//1.查询
			DB.Debug().First(&task, id)
			//2.删除
			DB.Debug().Delete(&task)
			//3.返回json
			ctx.JSON(http.StatusOK, gin.H{
				"code": 2000,
				"msg":  "删除成功",
				"data": nil,
			})
		})
	}

	ginServe.NoRoute(func(ctx *gin.Context) {
		ctx.HTML(http.StatusNotFound, "/templates/404.html", gin.H{})
	})
	ginServe.Run(":8000")
}
