package dao

import (
	"studygo2/CloudRestaurant/model"
	"studygo2/CloudRestaurant/tool"
)

type FoodCategoryDao struct {
	*tool.Orm
}

//实例化Dao
func NewFoodCategoryDao() *FoodCategoryDao {
	return &FoodCategoryDao{tool.DbEngine}
}

//从数据库查询美食列表
func (fcd *FoodCategoryDao) QueryCategories() ([]model.FoodCategory, error) {
	var categories []model.FoodCategory

	err := fcd.Engine.Find(&categories)
	if err != nil {
		return nil, err
	}
	return categories, nil
}
