package service

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go_gin/dao"
	"net/http"
)

// GetPageHandler get page
func GetPageHandler(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html", nil)
}

// GetAllTaskHandler get all task data
func GetAllTaskHandler(ctx *gin.Context) {
	var task []dao.Task
	tx := dao.DB.Debug().Find(&task)
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
		"data": task,
	})
}

// AddTaskHandler add one task data
func AddTaskHandler(ctx *gin.Context) {
	var data_ map[string]string
	data, _ := ctx.GetRawData()
	json.Unmarshal(data, &data_)
	content_ := data_["content"]
	task := dao.Task{Content: &content_}
	dao.DB.Debug().Create(&task)
	ctx.JSON(http.StatusOK, gin.H{
		"code": 2000,
		"msg":  "新增成功",
		"data": nil,
	})
}

// UpdataTaskHandler updata one task data
func UpdataTaskHandler(ctx *gin.Context) {
	var task dao.Task
	var body map[string]int
	bytes, _ := ctx.GetRawData()
	json.Unmarshal(bytes, &body)
	state := body["state"]
	id := ctx.Param("id")
	//1.查询
	tx := dao.DB.Debug().First(&task, id)
	if tx.RowsAffected == 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 2001,
			"msg":  "记录不存在",
			"data": nil,
		})
		return
	}
	//2.修改
	dao.DB.Debug().Model(&task).Update("Status", state)
	//3.返回json
	ctx.JSON(http.StatusOK, gin.H{
		"code": 2000,
		"msg":  "修改成功",
		"data": nil,
	})
}

// DelTaskHandler delect one task data
func DelTaskHandler(ctx *gin.Context) {
	var task dao.Task
	id := ctx.Param("id")
	//1.查询
	dao.DB.Debug().First(&task, id)
	//2.删除
	dao.DB.Debug().Delete(&task)
	//3.返回json
	ctx.JSON(http.StatusOK, gin.H{
		"code": 2000,
		"msg":  "删除成功",
		"data": nil,
	})
}
