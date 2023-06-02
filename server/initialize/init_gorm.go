package initialize

import (
	"AirGo/global"
	"AirGo/model"
	utils "AirGo/utils/encode_plugin"
	"errors"
	"gorm.io/driver/sqlite"
	"log"

	//"go-admin/initialize"
	//github.com/satori/go.uuid
	gormadapter "github.com/casbin/gorm-adapter/v3"
	uuid "github.com/satori/go.uuid"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// Gorm 初始化数据库并产生数据库全局变量
// Author SliverHorn
func Gorm() *gorm.DB {

	switch global.CONFIG.System.DbType {
	case "mysql":
		return GormMysql()
	case "sqlite":
		return GormSqlite()
	default:
		return GormMysql()
	}
}

// 初始化sqlite数据库
func GormSqlite() *gorm.DB {
	if db, err := gorm.Open(sqlite.Open(global.CONFIG.Sqlite.Path), &gorm.Config{
		SkipDefaultTransaction: true, //关闭事务，将获得大约 30%+ 性能提升
		NamingStrategy: schema.NamingStrategy{
			//TablePrefix: "gormv2_",
			SingularTable: true, //单数表名
		},
	}); err != nil {
		log.Println("gorm.Open error:", err)
		panic(err)
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(global.CONFIG.Mysql.MaxIdleConns)
		sqlDB.SetMaxOpenConns(global.CONFIG.Mysql.MaxOpenConns)
		return db
	}
}

// 初始化Mysql数据库
func GormMysql() *gorm.DB {
	mysqlConfig := mysql.Config{
		DSN:                       global.CONFIG.Mysql.Username + ":" + global.CONFIG.Mysql.Password + "@tcp(" + global.CONFIG.Mysql.Path + ":" + global.CONFIG.Mysql.Port + ")/" + global.CONFIG.Mysql.Dbname + "?" + global.CONFIG.Mysql.Config, // DSN data source name
		DefaultStringSize:         191,                                                                                                                                                                                                            // string 类型字段的默认长度
		SkipInitializeWithVersion: false,
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{
		SkipDefaultTransaction: true, //关闭事务，将获得大约 30%+ 性能提升
		NamingStrategy: schema.NamingStrategy{
			//TablePrefix: "gormv2_",
			SingularTable: true, //单数表名
		},
	}); err != nil {
		log.Println("gorm.Open error:", err)
		panic(err)
	} else {
		db.InstanceSet("gorm:table_options", "ENGINE="+global.CONFIG.Mysql.Engine)
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(global.CONFIG.Mysql.MaxIdleConns)
		sqlDB.SetMaxOpenConns(global.CONFIG.Mysql.MaxOpenConns)
		return db
	}
}

// RegisterTables 注册数据库表专用
func RegisterTables() {
	err := global.DB.AutoMigrate(
		// 用户表
		model.User{},
		//动态路由表
		model.DynamicRoute{},
		//角色表
		model.Role{},
		//商品
		model.Goods{},
		//订单
		model.Orders{},
		//流量统计表
		model.TrafficLog{},
		//主题
		model.Theme{},
		//系统设置参数
		model.Server{},
	)
	if err != nil {
		// global.LOG.Error("register table failed", zap.Error(err))
		//os.Exit(0)
		log.Println("table创建失败", err.Error())
		return
	}
	//global.LOG.Info("register table success")
	log.Println("table创建成功")
}

// 导入数据
func InsertInto(db *gorm.DB) error {
	uuid1 := uuid.NewV4()
	uuid2 := uuid.NewV4()
	sysUserData := []model.User{
		{
			UUID:     uuid1,
			UserName: "admin@oicq.com",
			Password: utils.BcryptEncode("admin"),
			NickName: "admin",
			//Avatar:   "https://qmplusimg.henrongyi.top/gva_header.jpg",
			Phone: "17611111111",
			//Email: "333333333@qq.com",
			SubscribeInfo: model.SubscribeInfo{
				GoodsID:      1,
				SubscribeUrl: "10010",
			},
		},
		{
			UUID:     uuid2,
			UserName: "测试1@oicq.com",
			Password: utils.BcryptEncode("123456"),
			NickName: "测试1",
			//Avatar:   "https://qmplusimg.henrongyi.top/gva_header.jpg",
			Phone: "17611111111",
			//Email: "333333333@qq.com",
		},
		{
			UUID:     uuid2,
			UserName: "测试2@oicq.com",
			Password: utils.BcryptEncode("123456"),
			NickName: "测试2",
			//Avatar:   "https://qmplusimg.henrongyi.top/gva_header.jpg",
			Phone: "",
			//Email: "",
		},
		{
			UUID:     uuid2,
			UserName: "测试3@oicq.com",
			Password: utils.BcryptEncode("123456"),
			NickName: "测试3",
			//Avatar:   "https://qmplusimg.henrongyi.top/gva_header.jpg",
			Phone: "",
			//Email: "",
		},
		{
			UUID:     uuid2,
			UserName: "测试4@oicq.com",
			Password: utils.BcryptEncode("123456"),
			NickName: "测试4",
			//Avatar:   "https://qmplusimg.henrongyi.top/gva_header.jpg",
			Phone: "",
			//Email: "",
		},
	}
	if err := db.Create(&sysUserData).Error; err != nil {
		return errors.New("db.Create(&userData) Error")
	}
	//插入sys_dynamic-router_data表
	DynamicRouteData := []model.DynamicRoute{
		{ParentID: 0, Path: "/home", Name: "home", Component: "/home/index.vue", Meta: model.Meta{Title: "首页", Icon: "iconfont icon-shouye"}},
		{ParentID: 0, Path: "/admin", Name: "admin", Component: "/layout/routerView/parent.vue", Meta: model.Meta{Title: "超级管理员", Icon: "iconfont icon-shouye_dongtaihui"}},
		{ParentID: 2, Path: "/admin/menu", Name: "adminMenu", Component: "/admin/menu/index.vue", Meta: model.Meta{Title: "菜单管理", Icon: "iconfont icon-caidan"}},
		{ParentID: 2, Path: "/admin/role", Name: "adminRole", Component: "/admin/role/index.vue", Meta: model.Meta{Title: "角色管理", Icon: "iconfont icon-icon-"}},
		{ParentID: 2, Path: "/admin/user", Name: "adminUser", Component: "/admin/user/index.vue", Meta: model.Meta{Title: "用户管理", Icon: "iconfont icon-gerenzhongxin"}},
		{ParentID: 2, Path: "/admin/order", Name: "adminOrder", Component: "/admin/order/index.vue", Meta: model.Meta{Title: "订单管理", Icon: "iconfont icon--chaifenhang"}},
		{ParentID: 2, Path: "/admin/node", Name: "adminNode", Component: "/admin/node/index.vue", Meta: model.Meta{Title: "节点管理", Icon: "iconfont icon-shuxingtu"}},
		{ParentID: 2, Path: "/admin/shop", Name: "adminShop", Component: "/admin/shop/index.vue", Meta: model.Meta{Title: "商品管理", Icon: "iconfont icon-zhongduancanshuchaxun"}},
		//{ParentID: 2, Path: "/admin/trafficRecord", Name: "trafficRecord", Component: "/admin/trafficRecord/index.vue", Meta: model.Meta{Title: "流量记录", Icon: "ele-ColdDrink"}},
		{ParentID: 2, Path: "/admin/system", Name: "system", Component: "/admin/system/index.vue", Meta: model.Meta{Title: "系统设置", Icon: "iconfont icon-xitongshezhi"}},
		{ParentID: 0, Path: "/shop", Name: "shop", Component: "/shop/index.vue", Meta: model.Meta{Title: "商店", Icon: "iconfont icon-zidingyibuju"}},
		{ParentID: 0, Path: "/myOrder", Name: "myOrder", Component: "/myOrder/index.vue", Meta: model.Meta{Title: "我的订单", Icon: "iconfont icon--chaifenhang"}},
		{ParentID: 0, Path: "/personal", Name: "personal", Component: "/personal/index.vue", Meta: model.Meta{Title: "个人信息", Icon: "iconfont icon-gerenzhongxin"}},
	}
	if err := db.Create(&DynamicRouteData).Error; err != nil {
		return errors.New("sys_dynamic-router_data表数据初始化失败!")
	}
	//插入user_role表
	sysRoleData := []model.Role{
		{ID: 1, RoleName: "admin", Description: "超级管理员"},
		{ID: 2, RoleName: "普通用户", Description: "普通用户"},
	}
	if err := db.Create(&sysRoleData).Error; err != nil {
		return errors.New("user_role表数据初始化失败!")
	}
	//插入user_and_role表
	userAndRoleData := []model.UserAndRole{
		{UserID: 1, RoleID: 1},
		{UserID: 1, RoleID: 2},
		{UserID: 2, RoleID: 2},
	}
	if err := db.Create(&userAndRoleData).Error; err != nil {
		return errors.New("user_and_role_data表数据初始化失败!")
	}
	//插入role_and_menu
	roleAndMenuData := []model.RoleAndMenu{
		{RoleID: 1, DynamicRouteID: 1},
		{RoleID: 1, DynamicRouteID: 2},
		{RoleID: 1, DynamicRouteID: 3},
		{RoleID: 1, DynamicRouteID: 4},
		{RoleID: 1, DynamicRouteID: 5},
		{RoleID: 1, DynamicRouteID: 6},
		{RoleID: 1, DynamicRouteID: 7},
		{RoleID: 1, DynamicRouteID: 8},
		{RoleID: 1, DynamicRouteID: 9},
		{RoleID: 1, DynamicRouteID: 10},
		{RoleID: 1, DynamicRouteID: 11},
		{RoleID: 1, DynamicRouteID: 12},

		{RoleID: 2, DynamicRouteID: 1},
		{RoleID: 2, DynamicRouteID: 8},
		{RoleID: 2, DynamicRouteID: 9},
	}
	if err := global.DB.Create(&roleAndMenuData).Error; err != nil {
		return errors.New("role_and_menu表数据初始化失败!")
	}
	//插入货物 goods
	goodsData := []model.Goods{
		{Subject: "10G|30天", TotalBandwidth: 10, ExpirationDate: 30, TotalAmount: "0.01"},
		{Subject: "20G|180天", TotalBandwidth: 20, ExpirationDate: 180, TotalAmount: "0.01"},
	}
	if err := global.DB.Create(&goodsData).Error; err != nil {
		return errors.New("goods表数据初始化失败!")
	}
	//插入node
	nodeData := []model.Node{
		{Name: "测试节点1", Address: "www.10010.com", Path: "/path", Port: "5566"},
		{Name: "测试节点2", Address: "www.10086.com", Path: "/path", Port: "5566"},
	}
	if err := global.DB.Create(&nodeData).Error; err != nil {
		return errors.New("node表数据初始化失败!")
	}
	//插入goods_and_nodes
	goodsAndNodesData := []model.GoodsAndNodes{
		{GoodsID: 1, NodeID: 1},
		{GoodsID: 1, NodeID: 2},
	}
	if err := global.DB.Create(&goodsAndNodesData).Error; err != nil {
		return errors.New("goods_and_nodes表数据初始化失败!")
	}
	// 插入casbin_rule
	casbinRuleData := []gormadapter.CasbinRule{
		{Ptype: "p", V0: "1", V1: "/public/getThemeConfig", V2: "GET"},

		{Ptype: "p", V0: "1", V1: "/user/register", V2: "POST"},
		{Ptype: "p", V0: "1", V1: "/user/login", V2: "POST"},
		{Ptype: "p", V0: "1", V1: "/user/getSub", V2: "GET"},
		{Ptype: "p", V0: "1", V1: "/user/changeSubHost", V2: "POST"},
		{Ptype: "p", V0: "1", V1: "/user/getUserInfo", V2: "GET"},
		{Ptype: "p", V0: "1", V1: "/user/changeUserPassword", V2: "POST"},
		{Ptype: "p", V0: "1", V1: "/user/resetUserPassword", V2: "POST"},
		{Ptype: "p", V0: "1", V1: "/user/resetSub", V2: "POST"},

		{Ptype: "p", V0: "1", V1: "/user/getUserList", V2: "POST"},
		{Ptype: "p", V0: "1", V1: "/user/newUser", V2: "POST"},
		{Ptype: "p", V0: "1", V1: "/user/updateUser", V2: "POST"},
		{Ptype: "p", V0: "1", V1: "/user/deleteUser", V2: "POST"},
		{Ptype: "p", V0: "1", V1: "/user/findUser", V2: "POST"},

		{Ptype: "p", V0: "1", V1: "/role/getRoleList", V2: "POST"},
		{Ptype: "p", V0: "1", V1: "/role/modifyRoleInfo", V2: "POST"},
		{Ptype: "p", V0: "1", V1: "/role/addRole", V2: "POST"},
		{Ptype: "p", V0: "1", V1: "/role/delRole", V2: "DELETE"},

		{Ptype: "p", V0: "1", V1: "/menu/getAllRouteList", V2: "GET"},
		{Ptype: "p", V0: "1", V1: "/menu/getAllRouteTree", V2: "GET"},
		{Ptype: "p", V0: "1", V1: "/menu/newDynamicRoute", V2: "POST"},
		{Ptype: "p", V0: "1", V1: "/menu/delDynamicRoute", V2: "POST"},
		{Ptype: "p", V0: "1", V1: "/menu/updateDynamicRoute", V2: "POST"},
		{Ptype: "p", V0: "1", V1: "/menu/findDynamicRoute", V2: "POST"},

		{Ptype: "p", V0: "1", V1: "/menu/getRouteList", V2: "GET"},
		{Ptype: "p", V0: "1", V1: "/menu/getRouteTree", V2: "GET"},

		//{Ptype: "p", V0: "1", V1: "/shop/alipayTradePreCreatePay", V2: "POST"},
		{Ptype: "p", V0: "1", V1: "/shop/preCreatePay", V2: "POST"},
		{Ptype: "p", V0: "1", V1: "/shop/purchase", V2: "POST"},
		{Ptype: "p", V0: "1", V1: "/shop/getAllEnabledGoods", V2: "GET"},
		{Ptype: "p", V0: "1", V1: "/shop/getAllGoods", V2: "GET"},
		{Ptype: "p", V0: "1", V1: "/shop/findGoods", V2: "POST"},
		{Ptype: "p", V0: "1", V1: "/shop/newGoods", V2: "POST"},
		{Ptype: "p", V0: "1", V1: "/shop/deleteGoods", V2: "POST"},
		{Ptype: "p", V0: "1", V1: "/shop/updateGoods", V2: "POST"},
		{Ptype: "p", V0: "1", V1: "/shop/alipayNotify", V2: "POST"},

		{Ptype: "p", V0: "1", V1: "/node/getAllNode", V2: "GET"},
		{Ptype: "p", V0: "1", V1: "/node/newNode", V2: "POST"},
		{Ptype: "p", V0: "1", V1: "/node/deleteNode", V2: "POST"},
		{Ptype: "p", V0: "1", V1: "/node/updateNode", V2: "POST"},
		{Ptype: "p", V0: "1", V1: "/node/getTraffic", V2: "POST"},

		{Ptype: "p", V0: "1", V1: "/mod_mu/nodes/:nodeID/info", V2: "GET"},
		{Ptype: "p", V0: "1", V1: "/mod_mu/users", V2: "GET"},
		{Ptype: "p", V0: "1", V1: "/mod_mu/users/traffic", V2: "POST"},
		{Ptype: "p", V0: "1", V1: "/mod_mu/users/aliveip", V2: "POST"},

		{Ptype: "p", V0: "1", V1: "/casbin/getPolicyByRoleIds", V2: "POST"},
		{Ptype: "p", V0: "1", V1: "/casbin/updateCasbinPolicy", V2: "POST"},
		{Ptype: "p", V0: "1", V1: "/casbin/updateCasbinPolicyNew", V2: "POST"},
		{Ptype: "p", V0: "1", V1: "/casbin/getPolicy", V2: "GET"},
		{Ptype: "p", V0: "1", V1: "/casbin/getAllPolicy", V2: "GET"},

		{Ptype: "p", V0: "1", V1: "/order/getOrderInfo", V2: "POST"},
		{Ptype: "p", V0: "1", V1: "/order/getAllOrder", V2: "POST"},
		{Ptype: "p", V0: "1", V1: "/order/getOrderByUserID", V2: "POST"},
		{Ptype: "p", V0: "1", V1: "/order/completedOrder", V2: "POST"},

		{Ptype: "p", V0: "1", V1: "/system/getThemeConfig", V2: "GET"},
		{Ptype: "p", V0: "1", V1: "/system/updateThemeConfig", V2: "POST"},
		{Ptype: "p", V0: "1", V1: "/system/getSetting", V2: "GET"},
		{Ptype: "p", V0: "1", V1: "/system/updateSetting", V2: "POST"},

		//普通用户
		{Ptype: "p", V0: "2", V1: "/public/getSetting", V2: "POST"},
		{Ptype: "p", V0: "2", V1: "/public/updateSetting", V2: "POST"},

		{Ptype: "p", V0: "2", V1: "/user/login", V2: "POST"},
		{Ptype: "p", V0: "2", V1: "/user/getSub", V2: "GET"},
		{Ptype: "p", V0: "2", V1: "/user/changeUserPassword", V2: "POST"},
		{Ptype: "p", V0: "2", V1: "/user/resetUserPassword", V2: "POST"},
		{Ptype: "p", V0: "2", V1: "/user/getUserInfo", V2: "GET"},
		{Ptype: "p", V0: "2", V1: "/user/resetSub", V2: "POST"},

		{Ptype: "p", V0: "2", V1: "/menu/getRouteList", V2: "GET"},
		{Ptype: "p", V0: "2", V1: "/menu/getRouteTree", V2: "GET"},

		{Ptype: "p", V0: "2", V1: "/order/getOrderInfo", V2: "POST"},

		{Ptype: "p", V0: "2", V1: "/shop/preCreatePay", V2: "POST"},
		{Ptype: "p", V0: "2", V1: "/shop/purchase", V2: "POST"},
		{Ptype: "p", V0: "2", V1: "/shop/getAllEnabledGoods", V2: "GET"},
		{Ptype: "p", V0: "2", V1: "/shop/findGoods", V2: "POST"},

		{Ptype: "p", V0: "2", V1: "/casbin/getPolicy", V2: "POST"},
	}
	if err := global.DB.Create(&casbinRuleData).Error; err != nil {
		return errors.New("casbin_rule表数据初始化失败!")
	}
	//主题配置
	themeData := model.Theme{
		ID: 1,
	}
	if err := global.DB.Create(&themeData).Error; err != nil {
		return errors.New("theme表数据初始化失败!")
	}

	//系统设置
	settingData := model.Server{
		ID: 1,
	}
	if err := global.DB.Create(&settingData).Error; err != nil {
		return errors.New("server表数据初始化失败!")
	}
	return nil
}
