package main

import (
	"github.com/fuzej/go_5_blog/controller"
	"github.com/fuzej/go_5_blog/dao/db"
	"github.com/gin-gonic/gin"
)

func main() {
	//1 加载数据库
	router := gin.Default()
	dns := `root:123456@tcp(127.0.0.1:3306)/?charset=utf8&parseTime=true`
	err := db.Init(dns)
	if err != nil {
		panic(err)
	}

	// 2 加载静态文件
	router.Static("/static/", "./static")
	router.LoadHTMLGlob("views/*")

	//3.
	router.GET("/", controller.IndexHandle)
	router.GET("/category", controller.CategoryList)
	router.Run(":8000")

}
