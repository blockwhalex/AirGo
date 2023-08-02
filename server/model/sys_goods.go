package model

import (
	"time"
)

type Goods struct {
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"-" gorm:"index"`
	ID        int        `json:"id" gorm:"primary_key"`
	//基础参数
	GoodsOrder  int    `json:"goods_order"` //排序
	Subject     string `json:"subject"   gorm:"comment:商品标题"`
	Des         string `json:"des"       gorm:"comment:描述"`
	TotalAmount string `json:"total_amount"  gorm:"comment:金额，单位为元，精确到小数点后两位"`

	CheckedNodes []int  `json:"checked_nodes" gorm:"-"` //前端套餐编辑时选中的节点
	Nodes        []Node `json:"nodes"         gorm:"many2many:goods_and_nodes"`
	Status       bool   `json:"status"        gorm:"default:true;comment:是否显示"`
	//订阅参数
	TotalBandwidth int64 `json:"total_bandwidth"` //总流量
	ExpirationDate int   `json:"expiration_date"` //有效期
	NodeConnector  int   `json:"node_connector"`  //可连接客户端数量
}

// 商品和节点 多对多 表
type GoodsAndNodes struct {
	GoodsID int
	NodeID  int
}
