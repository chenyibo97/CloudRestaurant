package controller

import (
	"github.com/gin-gonic/gin"
	"studygo2/CloudRestaurant/Service"
	"studygo2/CloudRestaurant/tool"
)

type ShopController struct {
}

func (sc *ShopController) Router(app *gin.Engine) {
	app.GET("/api/shops", sc.GetShopList)
	app.GET("/api/search_shops", sc.SearchShop)
}

func (sc *ShopController) GetShopList(ctx *gin.Context) {
	longtitude := ctx.Query("longtitude")
	latitude := ctx.Query("latitude")

	if longtitude == "" || latitude == "" {
		tool.Failed(ctx, "暂未获取到位置信息，请重试")
		longtitude = "116.34"
		latitude = "40.34"
		//return
	}
	shopService := Service.ShopService{}
	shopList := shopService.ShopList(longtitude, latitude)

	if len(shopList) == 0 {
		tool.Failed(ctx, "未获取到用户信息")
		return
	}
	for _, shop := range shopList {
		service := shopService.GetService(shop.Id)
		if len(service) == 0 {
			shop.Supports = nil
		} else {
			shop.Supports = &service
		}
	}
	tool.Sucess(ctx, shopList)

}

/*关键词搜索信息*/
func (sc *ShopController) SearchShop(ctx *gin.Context) {
	longtitude := ctx.Query("longtitude")
	latitude := ctx.Query("latitude")
	keyword := ctx.Query("keyword")
	if keyword == "" {
		tool.Failed(ctx, "重新输入商铺名称")
		return
	}
	shopService := Service.ShopService{}
	shopList := shopService.SerchShops(longtitude, latitude, keyword)
	if len(shopList) > 0 {
		tool.Sucess(ctx, shopList)
		return
	}
	tool.Failed(ctx, "未获取到商铺信息")
}
