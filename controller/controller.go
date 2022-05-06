package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"hujing.ltd/buttle/dao"
	"hujing.ltd/buttle/model"
)

func Index(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html", nil)
}

func PostTodo(ctx *gin.Context) {
	//从请求中把数据取出
	var todo model.Todo
	if err := ctx.ShouldBind(&todo); err != nil {
		log.Fatal(err)
	}
	//存入数据库
	err := dao.Create(&todo)
	//返回响应
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, todo)
	}

}

func LookupAllTodo(ctx *gin.Context) {
	todolist, err := dao.GetAll()
	if err != nil {
		log.Fatal("获取全部清单失败")
		ctx.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, todolist)
	}

}

func LookupOneTodo(ctx *gin.Context) {
	id := ctx.Params.ByName("id")
	ctx.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}

func ModifyOneTodo(ctx *gin.Context) {
	id := ctx.Params.ByName("id")
	err := dao.Update(id)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"error": err,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"status": "success",
		})
	}
}

func DeleteOneTodo(ctx *gin.Context) {
	id := ctx.Params.ByName("id")
	err := dao.Delete(id)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"error": err,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"status": "success",
		})
	}
}
