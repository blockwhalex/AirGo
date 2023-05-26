package service

import (
	"AirGo/global"
	"AirGo/model"
)

// 查询全部商品
func GetAllGoods() (*[]model.Goods, error) {
	var goodsArr []model.Goods
	err := global.DB.Find(&goodsArr).Error
	return &goodsArr, err

}

// 查询商品
func FindGoodsByGoodsID(goodsID int) (model.Goods, error) {
	var goods model.Goods
	err := global.DB.First(&goods, goodsID).Error
	return goods, err
}

// 查询商品
func FindGoodsByNodeID(nodeID int) ([]model.Goods, error) {
	//log.Println("nodeID:", nodeID)
	var node model.Node
	err := global.DB.Where("id = ?", nodeID).Preload("Goods").First(&node).Error
	if err != nil {
		return nil, err
	}
	//log.Println("FindGoodsByNodeID:", node.Goods)
	return node.Goods, nil

}

// 新建商品
func NewGoods(goods *model.Goods) error {
	//查询关联节点
	var nodeArr []model.Node
	global.DB.Where("id in ?", goods.CheckedNodes).Find(&nodeArr)
	goods.Nodes = nodeArr
	err := global.DB.Create(&goods).Error
	return err
}

// 删除商品
func DeleteGoods(goods *model.Goods) error {
	//删除关联
	err := global.DB.Debug().Model(&model.Goods{ID: goods.ID}).Association("Nodes").Replace(nil)
	if err != nil {
		//fmt.Println("删除商品参数err1", err)
		return err
	}
	err = global.DB.Debug().Where(&model.Goods{ID: goods.ID}).Delete(&model.Goods{}).Error
	return err

}

// 更新商品
func UpdateGoods(goods *model.Goods) error {
	//查询关联节点
	var nodeArr []model.Node
	global.DB.Where("id in ?", goods.CheckedNodes).Find(&nodeArr)
	goods.Nodes = nodeArr
	//fmt.Println("更新商品goods:", goods)
	//更新关联节点
	global.DB.Model(&goods).Association("Nodes").Replace(&goods.Nodes)
	// 更新商品
	err := global.DB.Model(&model.Goods{ID: goods.ID}).Save(&goods).Error
	return err
}
