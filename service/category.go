package service

import (
	"github.com/fuzej/go_5_blog/dao/db"
	"github.com/fuzej/go_5_blog/model"
)

// 获取所有分类
func GetALLCategoryList() (categoryList []*model.Category, err error) {
	categoryList, err = db.GetAllCategoryList()
	if err != nil {
		return
	}
	return
}
