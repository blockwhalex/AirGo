package model

import "time"

type Coupon struct {
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	DeletedAt    *time.Time `json:"-" gorm:"index"`
	ID           int        `json:"id" gorm:"primary_key"`
	Name         string
	DiscountRate float32
	Limit        int
	Goods        []Goods
}
