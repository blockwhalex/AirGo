package service

import (
	"AirGo/global"
	"AirGo/model"
	"fmt"
	"time"
)

// 创建系统订单
func CreateOrder(order *model.Orders) (*model.Orders, error) {
	err := global.DB.Debug().Create(&order).Error
	//直接获取新建的id
	fmt.Println(err)
	return order, err
}

// 更新数据库订单 by OutTradeNo
func UpdateOrder(order *model.Orders) error {
	//err := global.DB.Debug().Where("out_trade_no = ?", order.OutTradeNo).Updates(&order).Error
	//return err
	err := global.DB.Debug().Save(&order).Error
	return err
}

// 获取全部订单,分页获取
func GetAllOrder(orderParams *model.QueryParamsWithDate) (*model.OrdersWithTotal, error) {

	var startTime, endTime time.Time
	//时间格式转换
	if len(orderParams.Date) == 2 {
		startTime, _ = time.ParseInLocation("2006-01-02 15:04:05", orderParams.Date[0], time.Local)
		endTime, _ = time.ParseInLocation("2006-01-02 15:04:05", orderParams.Date[1], time.Local)
	} else {
		//默认前1个月数据
		endTime = time.Now().Local()
		startTime = endTime.AddDate(0, 0, -30)
		fmt.Println("默认前1个月数据", startTime, endTime)
	}
	var ordersWithTotal model.OrdersWithTotal
	var err error
	//预加载user
	if orderParams.Search != "" {
		err = global.DB.Model(&model.Orders{}).Count(&ordersWithTotal.Total).Where("out_trade_no = ? and created_at > ? and created_at < ?", orderParams.Search, startTime, endTime).Limit(orderParams.PageSize).Offset((orderParams.PageNum - 1) * orderParams.PageSize).Order("id desc").Find(&ordersWithTotal.OrderList).Error
	} else {
		err = global.DB.Model(&model.Orders{}).Count(&ordersWithTotal.Total).Where("created_at > ? and created_at < ?", startTime, endTime).Limit(orderParams.PageSize).Offset((orderParams.PageNum - 1) * orderParams.PageSize).Order("id desc").Find(&ordersWithTotal.OrderList).Error
	}
	return &ordersWithTotal, err

}

// 获取用户订单by user id,,默认显示最近10条
func GetOrderByUserID(userID int, orderParams *model.PaginationParams) (*[]model.Orders, error) {
	var orderArr []model.Orders
	return &orderArr, global.DB.Where("user_id = ?", userID).Limit(10).Order("id desc").Find(&orderArr).Error
}

// 获取用户订单
func GetOrderByOrderID(order *model.Orders) (*model.Orders, error) {
	var queryOrder model.Orders
	err := global.DB.Debug().Where(&model.Orders{OutTradeNo: order.OutTradeNo, UserID: order.UserID}).First(&queryOrder).Error
	return &queryOrder, err
}
