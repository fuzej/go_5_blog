package service

import (
	"fmt"
	"github.com/fuzej/go_5_blog/dao/db"
	"github.com/fuzej/go_5_blog/model"
)

func GetAllCategoryList() (categoryList []*model.Category, err error) {
	categoryList, err = db.GetAllCategoryList()
	if err != nil {
		fmt.Printf("get category list failed, err:%v\n", err)
		return
	}
	return
}
