package Service

import (
	"studygo2/CloudRestaurant/dao"
	"studygo2/CloudRestaurant/model"
)

type FoodCategoryService struct {
}

func NewFoodCategoryService() *FoodCategoryService {
	return &FoodCategoryService{}
}

/**
 * 获取美食类别
 */
func (fcs *FoodCategoryService) Categories() ([]model.FoodCategory, error) {
	fcd := dao.NewFoodCategoryDao()
	return fcd.QueryCategories()
}
