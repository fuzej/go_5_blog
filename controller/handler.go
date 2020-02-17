package controller

import (
	"github.com/fuzej/go_5_blog/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func IndexHandle(c *gin.Context) {
	//从service 取数据
	articleRecordList, err := service.GetArticleRecordList(0, 15)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
		return
	}

	//加载分类数据  gin本质上是一个map ——》 类似Attribute 键值对
	categoryList, err := service.GetALLCategoryList()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
	}
	c.HTML(http.StatusOK, "views/index.html", gin.H{
		"article_list":  articleRecordList,
		"category_list": categoryList,
	})

}

//点击分类云分类
func CategoryList(c *gin.Context) {
	categoryIdStr := c.Query("category_id")
	categoryId, err := strconv.ParseInt(categoryIdStr, 10, 64)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
		return
	}

	//拿到 id ,获取文章列表

}
