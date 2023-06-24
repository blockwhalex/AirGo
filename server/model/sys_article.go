package model

import "time"

type Article struct {
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"-" gorm:"index"`
	ID        int        `json:"id"      gorm:"primary_key"`

	Title        string `json:"title"        gorm:"comment:文字标题"`
	Introduction string `json:"introduction" gorm:"comment:文章简介"`
	Content      string `json:"content"      gorm:"comment:文章内容;size:3000"`
}

type ArticleWithTotal struct {
	Total       int64     `json:"total"`
	ArticleList []Article `json:"article_list"`
}
