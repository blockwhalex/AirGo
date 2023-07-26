package service

import (
	"AirGo/global"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"strconv"
	"time"

	"AirGo/model"
	encode_plugin "AirGo/utils/encode_plugin"
	"errors"
	uuid "github.com/satori/go.uuid"
)

// 注册
func Register(u *model.User) error {
	//判断是否存在
	var user model.User
	err := global.DB.Where(&model.User{UserName: u.UserName}).First(&user).Error
	if err == nil {
		return errors.New("用户已存在")
	} else if err == gorm.ErrRecordNotFound {
		var newUser = model.User{
			UUID:           uuid.NewV4(),
			UserName:       u.UserName,
			NickName:       u.UserName,
			Password:       encode_plugin.BcryptEncode(u.Password),
			RoleGroup:      []model.Role{{ID: 2}}, //默认角色
			InvitationCode: encode_plugin.RandomString(8),
			ReferrerCode:   u.ReferrerCode,
		}
		return CreateUser(NewUserSubscribe(&newUser))
	} else {
		return err
	}
}

// 新注册用户分配套餐
func NewUserSubscribe(u *model.User) *model.User {
	//查询商品信息
	if global.Server.System.DefaultGoods == "" {
		return u
	}
	var goods = model.Goods{
		Subject: global.Server.System.DefaultGoods,
	}
	g, err := FindGoods(&goods)
	if err != nil {
		return u
	}
	// 处理用户订阅信息
	return HandleUserSubscribe(u, g)

}

// 用户登录
func Login(u *model.UserLogin) (*model.User, error) {
	var user model.User
	err := global.DB.Where("user_name = ?", u.UserName).First(&user).Error
	if err == gorm.ErrRecordNotFound {
		return nil, errors.New("用户不存在")
	} else if !user.Enable {
		return nil, errors.New("用户已冻结")
	} else {
		if err := encode_plugin.BcryptDecode(u.Password, user.Password); err != nil {
			return nil, errors.New("密码错误")
		}
		return &user, err
	}
	return &user, err
}

// 获取当前请求节点（根据 node_id 参数判断）可连接的用户
func FindUsersByGoods(goods *[]model.Goods) (*[]model.SSUsers, error) {
	var goodsArr []int
	for _, v := range *goods {
		goodsArr = append(goodsArr, v.ID)
	}
	var users []model.SSUsers
	//err := global.DB.Model(&model.User{}).Where("goods_id in (?)", goodsArr).Find(&users).Error
	err := global.DB.Model(&model.User{}).Where("goods_id in (?) and sub_status = ?", goodsArr, true).Find(&users).Error
	return &users, err
}

// 查询订单属于哪个用户
func FindUsersByOrderID(outTradeNo string) (*model.User, error) {
	var order model.Orders
	err := global.DB.Where("out_trade_no = ?", outTradeNo).Preload("User").Find(&order).Error
	return &order.User, err
}

// 查用户 by user_id
func FindUserByID(id int) (*model.User, error) {
	var u model.User
	err := global.DB.First(&u, id).Error
	return &u, err
}

// 查询用户 by user_name(邮箱)
func FindUserByEmail(u *model.User) (*model.User, error) {
	err := global.DB.Where("user_name = ?", u.UserName).First(&u).Error
	return u, err
}

// 更新用户订阅信息
func UpdateUserSubscribe(order *model.Orders) error {
	//查询商品信息
	goods, _ := FindGoodsByGoodsID(order.GoodsID)
	//查询订单属于哪个用户
	u, _ := FindUserByID(order.UserID)
	//构建用户订阅信息
	user := HandleUserSubscribe(u, goods)
	//更新用户订阅信息
	return global.DB.Save(&user).Error
}

// 处理用户订阅信息
func HandleUserSubscribe(u *model.User, goods *model.Goods) *model.User {
	u.SubscribeInfo.T = goods.TotalBandwidth * 1024 * 1024 * 1024 // TotalBandwidth单位：GB。总流量单位：B
	u.SubscribeInfo.U = 0                                         //用户已用流量清零 //如果用struct ,gorm 不会更新“零值”
	u.SubscribeInfo.D = 0                                         //用户已用流量清零 //如果用struct ,gorm 不会更新“零值”
	if u.SubscribeInfo.SubscribeUrl == "" {
		u.SubscribeInfo.SubscribeUrl = encode_plugin.RandomString(8) //随机字符串订阅url
	}

	u.SubscribeInfo.GoodsID = goods.ID //当前订购的套餐
	u.SubscribeInfo.SubStatus = true   //订阅状态
	t := time.Now().AddDate(0, 0, goods.ExpirationDate)
	u.SubscribeInfo.ExpiredAt = &t //过期时间
	if goods.NodeConnector != 0 {
		u.SubscribeInfo.NodeConnector = goods.NodeConnector //连接客户端数
	}
	return u
}

// 批量更新用户流量信息
func UpdateUserTrafficInfo(userArr []model.User, userIds []int) error {
	var userArrQuery []model.User
	err := global.DB.Where("id in ?", userIds).Select("id", "u", "d").Find(&userArrQuery).Error
	if err != nil {
		return err
	}
	for item, _ := range userArrQuery {
		userArrQuery[item].SubscribeInfo.U = userArrQuery[item].SubscribeInfo.U + userArr[item].SubscribeInfo.U
		userArrQuery[item].SubscribeInfo.D = userArrQuery[item].SubscribeInfo.D + userArr[item].SubscribeInfo.D
	}
	return global.DB.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{"u", "d"}),
	}).Create(&userArrQuery).Error

}

// 用户流量，有效期 检测任务
func UserExpiryCheck() error {
	//fmt.Println("开始用户流量，有效期 检测任务")
	return global.DB.Exec("update user set sub_status = 0 where expired_at < ? or ( u + d ) > t", time.Now()).Error
}

// 修改混淆
func ChangeSubHost(uID int, host string) error {
	var u model.User
	u.ID = uID
	u.SubscribeInfo.Host = host
	return global.DB.Updates(&u).Error
}

// 获取自身信息
func GetUserInfo(uID int) (*model.User, error) {
	var user model.User
	return &user, global.DB.First(&user, uID).Error
}

// 获取用户列表,分页
func GetUserlist(params *model.PaginationParams) (*model.UsersWithTotal, error) {
	var userArr model.UsersWithTotal
	var err error
	if params.Search != "" {
		err = global.DB.Model(&model.User{}).Where("user_name like ?", ("%" + params.Search + "%")).Count(&userArr.Total).Limit(params.PageSize).Offset((params.PageNum - 1) * params.PageSize).Preload("RoleGroup").Find(&userArr.UserList).Error
	} else {
		err = global.DB.Model(&model.User{}).Count(&userArr.Total).Limit(params.PageSize).Offset((params.PageNum - 1) * params.PageSize).Preload("RoleGroup").Find(&userArr.UserList).Error
	}
	return &userArr, err
}

// 更新用户信息
func UpdateUser(u *model.User) error {
	return global.DB.Updates(&u).Error
}
func SaveUser(u *model.User) error {
	return global.DB.Save(&u).Error
}

// 创建用户
func CreateUser(u *model.User) error {
	return global.DB.Create(&u).Error
}

// 删除用户
func DeleteUser(u *model.User) error {
	return global.DB.Delete(&u).Error
}

// 重置用户密码
func ResetUserPassword(u *model.User) error {
	return global.DB.Model(&model.User{}).Where("user_name = ?", u.UserName).Updates(&u).Error
}

// 处理推荐人返利
func ReferrerRebate(uID int, receiptAmount string) {
	u, err := FindUserByID(uID)
	if err != nil || u.ReferrerCode == "" { //error或者推荐人为空
		return
	}
	//
	var uu model.User
	global.DB.Where(&model.User{InvitationCode: u.ReferrerCode}).First(&uu)
	a, _ := strconv.ParseFloat(receiptAmount, 64)
	uu.Remain = (uu.Remain + a) * global.Server.System.RebateRate
	global.DB.Save(&uu)
}

// 处理用户余额
func RemainHandle(uid int, remain string) {
	remainFloat64, _ := strconv.ParseFloat(remain, 64)
	if remainFloat64 == 0 {
		return
	}
	user, _ := FindUserByID(uid)
	user.Remain = user.Remain - remainFloat64
	SaveUser(user)
}
