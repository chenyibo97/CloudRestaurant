package dao

import (
	"studygo2/CloudRestaurant/model"
	"studygo2/CloudRestaurant/tool"
)

type GoodDao struct {
	*tool.Orm
}

func NewGoodDao() *GoodDao {
	return &GoodDao{tool.DbEngine}
}

/**
 * 获取商家的食品列表
 */
func (gd *GoodDao) QueryFoods(shop_id int64) ([]model.Goods, error) {
	var goods []model.Goods
	err := gd.Orm.Where("shop_id= ?", shop_id).Find(&goods)
	if err != nil {
		return nil, err
	}
	return goods, nil
}
