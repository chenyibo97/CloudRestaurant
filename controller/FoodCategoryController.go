package controller

import (
	"github.com/gin-gonic/gin"
	"studygo2/CloudRestaurant/Service"
	"studygo2/CloudRestaurant/tool"
)

type FoodCategoryController struct {
}

func (fcc *FoodCategoryController) Router(engine *gin.Engine) {
	//美食类别
	engine.GET("/api/food_category", fcc.foodCategory)
}
func (fcc *FoodCategoryController) foodCategory(ctx *gin.Context) {
	fcs := Service.NewFoodCategoryService()
	categories, err := fcs.Categories()
	if err != nil {
		tool.Failed(ctx, "请求失败："+err.Error())
		return
	}
	//转换数据格式
	for _, category := range categories {
		if category.ImageUrl != "" {
			category.ImageUrl = "./uploadfile/" + category.ImageUrl
		}
	}
}
