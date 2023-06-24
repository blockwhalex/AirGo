package api

import (
	"AirGo/global"
	"AirGo/model"
	"AirGo/service"
	"AirGo/utils/response"
	"github.com/gin-gonic/gin"
)

// 获取全部文章
func GetArticle(ctx *gin.Context) {
	var params model.PaginationParams
	err := ctx.ShouldBind(&params)
	if err != nil {
		global.Logrus.Error("获取文章参数错误:", err)
		response.Fail("获取文章参数错误", nil, ctx)
		return
	}
	Articles, err := service.GetArticle(params)
	if err != nil {
		global.Logrus.Error("获取文章错误:", err)
		response.Fail("获取文章错误", nil, ctx)
		return
	}
	response.OK("获取文章成功", Articles, ctx)
}

// 新建文章
func NewArticle(ctx *gin.Context) {
	var article model.Article
	err := ctx.ShouldBind(&article)
	if err != nil {
		global.Logrus.Error("新建文章参数错误:", err)
		response.Fail("新建文章参数错误", nil, ctx)
		return
	}
	err = service.NewArticle(&article)
	if err != nil {
		global.Logrus.Error("新建文章错误:", err)
		response.Fail("新建文章错误", nil, ctx)
		return
	}
	response.OK("新建文章成功", nil, ctx)
}

// 删除文章
func DeleteArticle(ctx *gin.Context) {
	var Article model.Article
	err := ctx.ShouldBind(&Article)
	if err != nil {
		global.Logrus.Error("删除文章参数错误:", err)
		response.Fail("删除文章参数错误", nil, ctx)
		return
	}
	err = service.DeleteArticle(&Article)
	if err != nil {
		global.Logrus.Error("删除文章错误:", err)
		response.Fail("删除文章错误", nil, ctx)
		return
	}
	response.OK("删除文章成功", nil, ctx)
}

// 更新文章
func UpdateArticle(ctx *gin.Context) {
	var Article model.Article
	err := ctx.ShouldBind(&Article)
	if err != nil {
		global.Logrus.Error("更新文章参数错误:", err)
		response.Fail("更新文章参数错误", nil, ctx)
		return
	}
	err = service.UpdateArticle(&Article)
	if err != nil {
		global.Logrus.Error("更新文章错误:", err)
		response.Fail("更新文章错误", nil, ctx)
		return
	}
	response.OK("更新文章成功", nil, ctx)

}
