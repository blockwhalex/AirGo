package model

import "time"

type Gallery struct {
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
	DeletedAt  *time.Time `json:"-" gorm:"index"`
	ID         int        `json:"id" gorm:"primary_key"`
	Subject    string     `json:"subject"       gorm:"comment:图片标题"`
	PictureUrl string     `json:"picture_url"   gorm:"comment:图片链接"`
	UserID     int        `json:"user_id"       gorm:"comment:用户ID"`
}

type GalleryHeader struct {
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
	ID         string `json:"id"`
	Subject    string `json:"subject"`
	PictureUrl string `json:"picture_url"`
	UserID     string `json:"user_id"`
}

var GalleryHeaderItem = GalleryHeader{
	CreatedAt:  "创建日期",
	UpdatedAt:  "更新日期",
	ID:         "ID",
	Subject:    "图片标题",
	PictureUrl: "图片链接",
	UserID:     "用户ID",
}
