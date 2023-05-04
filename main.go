package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go_gin/dao"
	"gorm.io/gorm"
	"net/http"
	"os"
)

var DB *gorm.DB

func init() {
	DB = dao.Connect()
}

func main() {
	root, _ := os.Getwd()
	ginServe := gin.Default()
	ginServe.Static("/css", root+"/static/css")
	ginServe.Static("/fonts", root+"/static/fonts")
	ginServe.Static("/js", root+"/static/js")

	ginServe.LoadHTMLGlob("templates/*")

	ginServe.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", nil)
	})
	routerGroup := ginServe.Group("/api")
	{
		//查询
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

		//添加
		routerGroup.POST("/", func(ctx *gin.Context) {
			var data_ map[string]string
			data, _ := ctx.GetRawData()
			json.Unmarshal(data, &data_)
			content_ := data_["content"]
			task := dao.Task{Content: &content_}
			DB.Debug().Create(&task)
			ctx.JSON(http.StatusOK, gin.H{
				"code": 2000,
				"msg":  "新增成功",
				"data": nil,
			})
		})

		//修改
		routerGroup.PUT("/:id", func(ctx *gin.Context) {
			var task dao.Task
			var body map[string]int
			//ctx.BindJSON(&task)
			bytes, _ := ctx.GetRawData()
			json.Unmarshal(bytes, &body)
			state := body["state"]
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

		//删除
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
		ctx.HTML(http.StatusNotFound, "404.html", gin.H{})
	})
	ginServe.Run(":8000")
}
