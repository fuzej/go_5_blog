package main

import (
	"github.com/DeanThompson/ginpprof"
	"github.com/fuzej/go_5_blog/controller"
	"github.com/fuzej/go_5_blog/dao/db"
	"github.com/gin-gonic/gin"
)

func main() {
	//1 加载数据库
	router := gin.Default()
	dns := `root:123456@tcp(127.0.0.1:3306)/blogger?charset=utf8&parseTime=true`
	err := db.Init(dns)
	if err != nil {
		panic(err)
	}
	ginpprof.Wrapper(router)
	router.Static("/static/", "./static")
	router.LoadHTMLGlob("views/*")
	router.GET("/", controller.IndexHandle)
	router.GET("/category/", controller.CategoryList)
	router.GET("/article/new/", controller.NewArticle)
	router.POST("/article/submit/", controller.ArticleSubmit)
	router.GET("/article/detail/", controller.ArticleDetail)
	router.POST("/upload/file/", controller.UploadFile)
	router.GET("/leave/new/", controller.LeaveNew)
	router.GET("/about/me/", controller.AboutMe)
	router.POST("/comment/submit/", controller.CommentSubmit)
	router.POST("/leave/submit/", controller.LeaveSubmit)
	_ = router.Run(":8000")

}
