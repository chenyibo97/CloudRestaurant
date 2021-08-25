package dao

import (
	"fmt"
	"studygo2/CloudRestaurant/model"
	"studygo2/CloudRestaurant/tool"
)

const DEFAULT_RANGE = 5

type ShopDao struct {
	*tool.Orm
}

func NewShopDao() *ShopDao {
	return &ShopDao{tool.DbEngine}
}
func (shopDao *ShopDao) QueryServiceByShopId(shopId int64) []model.Service {
	var service []model.Service
	err := shopDao.Table("service").Join("INNER", "shop_service", " service.id = shop_service.service_id and shop_service.shop_id = ? ", shopId).Find(&service)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	return service
}

/*func(shopdao *ShopDao)QueryShops(longitude,latitude float64)[]model.Shop{
	var shops []model.Shop
	err := shopdao.Engine.Where("longitude> ? and longitude < ? and latitude >? and latitude <? ",
		longitude-DEFALUT_RANGE, longitude+DEFALUT_RANGE,
		latitude-DEFALUT_RANGE, longitude+DEFALUT_RANGE).Find(&shops)
	if err!=nil{
		fmt.Println("查询商店数据库失败:",err)
		return nil
	}
	return shops

}*/
func (shopDao *ShopDao) QueryShops(longtitude, latitude float64, keyword string) []model.Shop {
	var shops []model.Shop

	if keyword == "" {
		err := shopDao.Where(" longitude > ? and longitude < ?  and latitude > ? and latitude < ? and status = 1", longtitude-DEFAULT_RANGE, longtitude+DEFAULT_RANGE, latitude-DEFAULT_RANGE, latitude+DEFAULT_RANGE).Find(&shops)
		if err != nil {
			fmt.Println(err.Error())
			return nil
		}
	} else {
		err := shopDao.Where(" longitude > ? and longitude < ?  and latitude > ? and latitude < ?  and name like ? and status = 1", longtitude-DEFAULT_RANGE, longtitude+DEFAULT_RANGE, latitude-DEFAULT_RANGE, latitude+DEFAULT_RANGE, keyword).Find(&shops)
		if err != nil {
			fmt.Println(err.Error())
			return nil
		}
	}
	return shops
}
