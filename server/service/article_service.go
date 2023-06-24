package service

import (
	"AirGo/global"
	"AirGo/model"
)

func GetArticle(params model.PaginationParams) (*model.ArticleWithTotal, error) {
	var articleArr model.ArticleWithTotal
	if params.Search != "" {
		err := global.DB.Model(&model.Article{}).Count(&articleArr.Total).Where("title LIKE ?", "%"+params.Search+"%").Limit(params.PageSize).Offset((params.PageNum - 1) * params.PageSize).Order("id desc").Find(&articleArr.ArticleList).Error
		if err != nil {
			global.Logrus.Error("查询文章error:", err.Error())
			return &model.ArticleWithTotal{}, err
		}
	} else {
		err := global.DB.Debug().Model(&model.Article{}).Count(&articleArr.Total).Limit(params.PageSize).Offset((params.PageNum - 1) * params.PageSize).Order("id desc").Find(&articleArr.ArticleList).Error
		if err != nil {
			global.Logrus.Error("查询文章error:", err.Error())
			return &model.ArticleWithTotal{}, err
		}
	}
	return &articleArr, nil
}
func NewArticle(article *model.Article) error {
	return global.DB.Create(&article).Error
}
func DeleteArticle(article *model.Article) error {
	return global.DB.Delete(&article).Error
}
func UpdateArticle(article *model.Article) error {
	return global.DB.Save(&article).Error
}
