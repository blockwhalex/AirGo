package api

import (
	"AirGo/global"
	"AirGo/model"
	"AirGo/service"
	"AirGo/utils/response"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

// 上传图片链接
func NewPictureUrl(ctx *gin.Context) {
	picUrl := ctx.Query("picUrl")
	subject := ctx.Query("subject")
	fmt.Println("url sub:", picUrl, subject)

	if picUrl == "" {
		response.Fail("上传图片链接参数错误", nil, ctx)
		return
	}
	if subject == "" {
		subject = time.Now().Format("2006-01-02 15:03:04")
	}
	uID, _ := ctx.Get("uID")
	uIDInt := uID.(int)
	err := service.NewPictureUrl(uIDInt, picUrl, subject)
	if err != nil {
		response.Fail("上传图片链接错误", nil, ctx)
		return
	}
	response.OK("上传图片链接成功", nil, ctx)
}
func GetPictureList(ctx *gin.Context) {
	var params model.PaginationParams
	err := ctx.ShouldBind(&params)
	if err != nil {
		global.Logrus.Error("获取图片列表错误：", err.Error())
		response.Fail("获取图片列表错误："+err.Error(), nil, ctx)
		return
	}
	fmt.Println("params:", params)
	picList, err := service.GetPictureList(&params)
	if err != nil {
		global.Logrus.Error("获取图片列表错误：", err.Error())
		response.Fail("获取图片列表错误："+err.Error(), nil, ctx)
		return
	}
	response.OK("获取图片列表成功：", picList, ctx)
}
