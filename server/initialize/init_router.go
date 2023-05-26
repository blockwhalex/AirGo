package initialize

import (
	"AirGo/api"
	"AirGo/global"
	"AirGo/middleware"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 初始化总路由
func InitRouter() {
	Router := gin.Default()
	Router.Use(middleware.Cors()) //不开启跨域验证码出错
	RouterGroup := Router.Group("/")

	//公共路由
	publicRouter := RouterGroup.Group("public")
	{
		publicRouter.POST("getEmailCode", api.GetMailCode)
		publicRouter.GET("getThemeConfig", api.GetThemeConfig) //获取主题配置
	}

	//user
	userRouter := RouterGroup.Group("user").Use(middleware.ParseJwt(), middleware.CasbinMiddleware())
	{
		userRouter.POST("changeSubHost", api.ChangeSubHost) //修改混淆
		userRouter.GET("getUserInfo", api.GetUserInfo)      //获取自身信息

		userRouter.POST("getUserList", api.GetUserlist)               //获取用户列表
		userRouter.POST("newUser", api.NewUser)                       //新建用户
		userRouter.POST("updateUser", api.UpdateUser)                 //修改用户
		userRouter.POST("deleteUser", api.DeleteUser)                 //删除用户
		userRouter.POST("changeUserPassword", api.ChangeUserPassword) //修改密码
		//	userRouter.POST("findUser", api.Finduser)       //查询用户
	}
	userRouterNoVerify := RouterGroup.Group("user")
	{
		userRouterNoVerify.GET("ping", func(c *gin.Context) { c.JSON(200, "success") })
		userRouterNoVerify.POST("register", api.Register)                   //用户注册
		userRouterNoVerify.POST("login", api.Login)                         //用户登录
		userRouterNoVerify.GET("getSub", api.GetSub)                        //获取订阅
		userRouterNoVerify.POST("resetUserPassword", api.ResetUserPassword) //重置密码
	}

	//菜单
	menuRouter := RouterGroup.Group("menu").Use(middleware.ParseJwt())
	{
		menuRouter.GET("getAllRouteList", api.GetAllRouteList)        //获取全部角色动态路由
		menuRouter.GET("getAllRouteTree", api.GetAllRouteTree)        //获取全部动态路由节点树
		menuRouter.POST("newDynamicRoute", api.NewDynamicRoute)       //新建动态路由
		menuRouter.POST("delDynamicRoute", api.DelDynamicRoute)       //删除动态路由
		menuRouter.POST("updateDynamicRoute", api.UpdateDynamicRoute) //修改动态路由
		menuRouter.POST("findDynamicRoute", api.FindDynamicRoute)     //查询单条动态路由 by meta.title

		menuRouter.GET("getRouteList", api.GetRouteList) //获取当前角色动态路由
		menuRouter.GET("getRouteTree", api.GetRouteTree) //获取当前角色动态路由节点树
	}

	//角色
	roleRouter := RouterGroup.Group("role")
	{
		roleRouter.POST("getAllRoleList", api.GetRoleList)    //获取role list
		roleRouter.POST("modifyRoleInfo", api.ModifyRoleInfo) //更新role 信息
		roleRouter.POST("addRole", api.AddRole)               //更新role 信息
		roleRouter.DELETE("delRole", api.DelRole)             //删除role
	}

	//系统设置
	systemRouter := RouterGroup.Group("system")
	systemRouter.Use(middleware.ParseJwt(), middleware.CasbinMiddleware())
	{
		systemRouter.GET("getThemeConfig", api.GetThemeConfig)        //获取主题配置
		systemRouter.POST("updateThemeConfig", api.UpdateThemeConfig) //设置主题配置
		systemRouter.GET("getSetting", api.GetSetting)
		systemRouter.POST("updateSetting", api.UpdateSetting)

	}

	//节点
	nodeRouter := RouterGroup.Group("node")
	nodeRouter.Use(middleware.ParseJwt(), middleware.CasbinMiddleware())
	{
		nodeRouter.POST("getNodeByName", api.GetNodeByName) //根据node name 模糊查询节点
		nodeRouter.GET("getAllNode", api.GetAllNode)        //查询全部节点
		nodeRouter.POST("newNode", api.NewNode)             //新建节点
		nodeRouter.POST("deleteNode", api.DeleteNode)       //删除节点
		nodeRouter.POST("updateNode", api.UpdateNode)       //更新节点
		nodeRouter.POST("getTraffic", api.GetNodeTraffic)   //获取节点流量统计
	}

	//sspqnel
	sspanelRouter := RouterGroup.Group("mod_mu")
	//sspanelRouter.Use(middleware.ParseJwt())
	{
		sspanelRouter.GET("nodes/:nodeID/info", api.SSNodeInfo) //获取节点信息
		sspanelRouter.GET("users", api.SSUsers)                 //获取当前节点可连接的用户
		sspanelRouter.POST("users/traffic", api.SSUsersTraffic) //上报用户的流量使用情况
		sspanelRouter.POST("users/aliveip", api.SSUsersAliveIP) //上报用户的当前在线IP
	}

	//商店
	shopRouter := RouterGroup.Group("shop")
	shopRouter.Use(middleware.ParseJwt(), middleware.CasbinMiddleware())
	{
		shopRouter.POST("preCreatePay", api.PreCreatePay) //alipay,统一收单线下交易预创建
		shopRouter.POST("purchase", api.Purchase)         //支付

		shopRouter.GET("getAllGoods", api.GetAllGoods)  //查询全部商品
		shopRouter.POST("findGoods", api.FindGoods)     //查询商品
		shopRouter.POST("newGoods", api.NewGoods)       //新建商品
		shopRouter.POST("deleteGoods", api.DeleteGoods) //删除商品
		shopRouter.POST("updateGoods", api.UpdateGoods) //更新商品
	}
	shopRouterNoVerify := RouterGroup.Group("/shop")
	{
		shopRouterNoVerify.POST("alipayNotify", api.AlipayNotify) //异步验证支付结果
	}
	//订单
	orderRouter := RouterGroup.Group("order")
	orderRouter.Use(middleware.ParseJwt(), middleware.CasbinMiddleware())
	{
		orderRouter.POST("getAllOrder", api.GetAllOrder)           //获取全部订单，分页获取
		orderRouter.POST("getOrderByUserID", api.GetOrderByUserID) //获取订单，分页获取
	}

	//casbin
	casbinRouter := RouterGroup.Group("casbin")
	casbinRouter.Use(middleware.ParseJwt(), middleware.CasbinMiddleware())
	//casbinRouter.Use(middleware.ParseJwt())
	{

		casbinRouter.GET("getAllPolicy", api.GetAllPolicy)                    //获取全部权限
		casbinRouter.GET("getPolicy", api.GetPolicy)                          //获取用户自身权限
		casbinRouter.POST("getPolicyByRoleIds", api.GetPolicyByRoleIds)       //获取用户权限ByRoleIds
		casbinRouter.POST("updateCasbinPolicy", api.UpdateCasbinPolicy)       //更新casbin权限
		casbinRouter.POST("updateCasbinPolicyNew", api.UpdateCasbinPolicyNew) //更新casbin权限
	}
	Router.Run(":" + strconv.Itoa(global.CONFIG.System.Port))
}
