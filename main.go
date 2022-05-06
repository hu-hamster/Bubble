package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"hujing.ltd/buttle/controller"
	"hujing.ltd/buttle/dao"
)

func init() {
	//添加viper配置
	dir, err := os.Getwd()
	if err != nil {
		log.Panic("获取项目路径失败")
	}
	viper.AddConfigPath(dir + "/config")
	viper.SetConfigType("yml")
	viper.SetConfigName("application")
	if err := viper.ReadInConfig(); err != nil {
		log.Panic(err)
	}
	//初始化数据库
	dao.InitDB()

}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/static", "./static")
	r.GET("/", func(ctx *gin.Context) {
		ctx.HTML(200, "index.html", nil)
	})

	bubble := r.Group("v1")
	bubble.POST("/todo", controller.PostTodo)
	bubble.GET("/todo", controller.LookupAllTodo)
	bubble.GET("/todo/:id", controller.LookupOneTodo)
	bubble.PUT("/todo/:id", controller.ModifyOneTodo)
	bubble.DELETE("/todo/:id", controller.DeleteOneTodo)
	log.Panic(r.Run(":8888"))
}
