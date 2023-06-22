package api

import (
	"AirGo/global"
	"AirGo/model"
	"AirGo/service"
	"AirGo/utils/response"
	"github.com/gin-gonic/gin"
)

func ReportFind(ctx *gin.Context) {

}

// 获取数据库的所有数据库名
func GetDB(ctx *gin.Context) {
	//var database model.DbNameAndTable
	//err := ctx.ShouldBind(&database)
	//if err != nil {
	//	global.Logrus.Error("获取数据库的所有数据库名 error:", err.Error())
	//	response.Fail("获取数据库的所有数据库名 error:"+err.Error(), nil, ctx)
	//	return
	//}
	res, err := service.GetDB()
	if err != nil {
		global.Logrus.Error("获取数据库的所有数据库名 error:", err.Error())
		response.Fail("获取数据库的所有数据库名 error:"+err.Error(), nil, ctx)
		return
	}
	response.OK("获取数据库的所有数据库名成功", res, ctx)
}

// 获取数据库的所有表名
func GetTables(ctx *gin.Context) {
	var dbName model.DbName
	err := ctx.ShouldBind(&dbName)
	if err != nil {
		global.Logrus.Error("获取数据库的所有表名 error:", err.Error())
		response.Fail("获取数据库的所有表名 error:"+err.Error(), nil, ctx)
		return
	}
	if dbName.Database == "" {
		response.Fail("数据库名为空", nil, ctx)
		return
	}
	res, err := service.GetTables(dbName.Database)
	if err != nil {
		global.Logrus.Error("获取数据库的所有表名 error:", err.Error())
		response.Fail("获取数据库的所有表名 error:"+err.Error(), nil, ctx)
		return
	}
	response.OK("获取数据库的所有表名成功", res, ctx)
}

// 获取 数据库 数据表的所有字段名,类型值
func GetColumn(ctx *gin.Context) {
	//body, err := ioutil.ReadAll(ctx.Request.Body)
	//fmt.Println("body:", string(body), err)
	var dbNameAndTable model.DbNameAndTable
	err := ctx.ShouldBind(&dbNameAndTable)
	if err != nil {
		global.Logrus.Error("获取数据表所有字段名, error:", err.Error())
		response.Fail("获取数据表所有字段名, error:"+err.Error(), nil, ctx)
		return
	}
	if dbNameAndTable.Database == "" {
		if global.Config.SystemParams.DbType == "mysql" {
			dbNameAndTable.Database = global.Config.Mysql.Dbname
		} else {
			dbNameAndTable.Database = global.Config.Sqlite.Path
		}
	}
	//fmt.Println("所有字段名:", dbNameAndTable)
	res, err := service.GetColumn(dbNameAndTable.Database, dbNameAndTable.TableName)
	if err != nil {
		global.Logrus.Error("获取数据表所有字段名, error:", err.Error())
		response.Fail("获取数据表所有字段名, error:"+err.Error(), nil, ctx)
		return
	}
	response.OK("获取数据库所有字段名成功", res, ctx)
}

// 获取报表
func ReportSubmit(ctx *gin.Context) {
	//body, err := ioutil.ReadAll(ctx.Request.Body)
	//fmt.Println("body:", string(body), err)

	var fieldParams model.FieldParams
	err := ctx.ShouldBind(&fieldParams)
	if err != nil {
		global.Logrus.Error("获取报表, error:", err.Error())
		response.Fail("获取报表,error:"+err.Error(), nil, ctx)
		return
	}
	//fmt.Println("获取报表:", fieldParams)
	res, total, err := service.GetReport(fieldParams)
	if err != nil {
		global.Logrus.Error("获取报表,error:", err.Error())
		response.Fail("获取报表, error:"+err.Error(), nil, ctx)
		return
	}
	var tableHeader interface{}
	switch fieldParams.TableName {
	case "user":
		tableHeader = model.UserHeaderItem
	case "orders":
		tableHeader = model.OrdersHeaderItem
	case "gallery":
		tableHeader = model.GalleryHeaderItem
	default:
	}
	response.OK("获取报表成功", gin.H{
		"table_header": tableHeader,
		"table_data":   res,
		"total":        total,
	}, ctx)

}
