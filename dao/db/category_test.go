package db

import (
	"testing"
)

func init() {
	//初始化 数据库
	dns := `root:123456@tcp(127.0.0.1:3306)/blogger?charset=utf8&parseTime=true`
	err := Init(dns)
	if err != nil {
		panic(err.Error())
	}
}

func TestGetCategoryById(t *testing.T) {
	category, err := GetCategoryById(1)
	if err != nil {
		panic(err)
	}
	t.Logf("category:%#v", category)
}

func TestGetCategoryList(t *testing.T) {
	//切片的使用
	var categoryIds []int64
	categoryIds = append(categoryIds, 1, 2, 3)
	categoryList, err := GetCategoryList(categoryIds)
	if err != nil {
		panic(err)
	}
	for _, v := range categoryList {
		t.Logf("id:%d category:%#v\n", v.CategoryId, v)
	}
}
func TestGetAllCategoryList(t *testing.T) {
	categoryList, err := GetAllCategoryList()
	if err != nil {
		panic(err)
	}
	for _, v := range categoryList {
		t.Logf("id:%d category:%#v\n", v.CategoryId, v)
	}
}
