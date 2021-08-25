package Service

import (
	"fmt"
	"strconv"
	"studygo2/CloudRestaurant/dao"
	"studygo2/CloudRestaurant/model"
)

type ShopService struct {
}

func (s *ShopService) GetService(shopId int64) []model.Service {
	shopDao := dao.NewShopDao()
	return shopDao.QueryServiceByShopId(shopId)
}

/*查询数据库商店列表
 */
func (s *ShopService) ShopList(long, lat string) []model.Shop {
	longtitude, err := strconv.ParseFloat(long, 10)
	if err != nil {
		fmt.Println("解析经纬度失败")
		return nil
	}
	latitude, err := strconv.ParseFloat(lat, 10)
	if err != nil {
		fmt.Println("解析经纬度失败")
		return nil
	}
	shopDao := dao.NewShopDao()
	return shopDao.QueryShops(longtitude, latitude, "")
}

func (s *ShopService) SerchShops(long, lat, keyword string) []model.Shop {
	longtitude, err := strconv.ParseFloat(long, 10)
	if err != nil {
		fmt.Println("解析经纬度失败")
		return nil
	}
	latitude, err := strconv.ParseFloat(lat, 10)
	if err != nil {
		fmt.Println("解析经纬度失败")
		return nil
	}
	shopDao := dao.NewShopDao()
	return shopDao.QueryShops(longtitude, latitude, keyword)
}
