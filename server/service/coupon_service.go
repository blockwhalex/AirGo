package service

import (
	"AirGo/global"
	"AirGo/model"
	"errors"
	"gorm.io/gorm"
	"time"
)

func NewCoupon(coupon model.Coupon) error {
	return global.DB.Debug().Create(&coupon).Error
}
func DeleteCoupon(coupon model.Coupon) error {
	return global.DB.Delete(&coupon).Error
}
func UpdateCoupon(coupon model.Coupon) error {
	return global.DB.Save(&coupon).Error
}
func GetCoupon() (*[]model.Coupon, error) {
	var couponArr []model.Coupon
	err := global.DB.Order("id desc").Find(&couponArr).Error
	return &couponArr, err
}

func VerifyCoupon(coupon string, userId int) (model.Coupon, error) {
	var c model.Coupon
	err := global.DB.Where(&model.Coupon{Name: coupon}).First(&c).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return model.Coupon{}, errors.New("优惠码不存在")
		} else {
			return model.Coupon{}, errors.New("优惠码错误")
		}
	}
	if time.Now().After(c.ExpiredAt) {
		return model.Coupon{}, errors.New("优惠码已过期")
	}
	//判断使用次数
	orderArr, err := GetOrderByCouponID(userId, c.ID)
	if err != nil {
		return model.Coupon{}, errors.New("优惠码错误")
	}
	if len(orderArr) >= c.Limit {
		return model.Coupon{}, errors.New("优惠码次数用尽")
	}
	//返回
	return c, nil
}
