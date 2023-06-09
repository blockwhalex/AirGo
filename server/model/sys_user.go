package model

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// gorm 创建/更新时间追踪（纳秒、毫秒、秒、Time）
// CreatedAt time.Time // 在创建时，如果该字段值为零值，则使用当前时间填充
// UpdatedAt int       // 在创建时该字段值为零值或者在更新时，使用当前时间戳秒数填充
// Updated   int64 `gorm:"autoUpdateTime:nano"` // 使用时间戳纳秒数填充更新时间
// Updated   int64 `gorm:"autoUpdateTime:milli"` // 使用时间戳毫秒数填充更新时间
// Created   int64 `gorm:"autoCreateTime"`      // 使用时间戳秒数填充创建时间
// 用户
type User struct {
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"-" gorm:"index"`

	ID       int       `json:"id"      gorm:"primary_key"`
	UUID     uuid.UUID `json:"uuid"         gorm:"comment:用户UUID"`                                                                        // 用户UUID
	UserName string    `json:"user_name"    gorm:"comment:用户登录名"`                                                                         // 用户登录名,邮箱
	Password string    `json:"password"     gorm:"comment:用户登录密码"`                                                                        // 用户登录密码
	NickName string    `json:"nick_name"    gorm:"default:系统用户;comment:用户昵称"`                                                             // 用户昵称
	Avatar   string    `json:"avatar"       gorm:"default:https://telegraph-image.pages.dev/file/28f40afe1021a81434bfa.jpg;comment:用户头像"` // 用户头像
	Phone    string    `json:"phone"        gorm:"comment:用户手机号"`                                                                         // 用户手机号
	//Email  string `json:"email"       gorm:"comment:用户邮箱"`                                                                                             // 用户邮箱
	Enable bool `json:"enable"      gorm:"default:true;comment:用户是否被冻结 true正常 false冻结"` //用户是否被冻结 1正常 2冻结
	//角色组
	RoleGroup []Role `json:"role_group" gorm:"many2many:user_and_role;"` //多对多
	//订单
	Orders []Orders `json:"orders" gorm:"foreignKey:UserID;references:ID"` //has many
	//附加订阅信息
	SubscribeInfo SubscribeInfo `json:"subscribe_info" gorm:"embedded;comment:附加订阅信息"`
}

// 附加订阅信息
type SubscribeInfo struct {
	Host           string     `json:"host"              gorm:"comment:用户混淆"`
	ClientIP       string     `json:"client_ip"         gorm:"comment:用户的当前在线IP"`              //用户的当前在线IP
	SubStatus      bool       `json:"sub_status"        gorm:"comment:订阅是否有效"`                 //订阅是否有效
	SubscribeUrl   string     `json:"subscribe_url"     gorm:"comment:订阅链接"`                   //订阅链接
	GoodsID        int        `json:"goods_id"          gorm:"comment:商品ID"`                   //商品ID
	ExpiredAt      *time.Time `json:"expired_at"        gorm:"comment:过期时间"`                   //过期时间
	T              int        `json:"t"                 gorm:"comment:总流量（Byte）"`              //总流量（Byte）
	U              int        `json:"u"                 gorm:"comment:上行流量"`                   //上行流量（Byte）
	D              int        `json:"d"                 gorm:"comment:下行流量"`                   //下行流量（Byte）
	ResetDay       int        `json:"reset_day"         gorm:"comment:流量重置日"`                  //流量重置日
	NodeSpeedLimit int        `json:"node_speedlimit"   gorm:"default:0;comment:限速Mbps（Mbps）"` //限速Mbps（Mbps）
	NodeConnector  int        `json:"node_connector"    gorm:"default:3;comment:连接客户端数"`       //连接客户端数
	//IsMultiUser    int    `json:"is_multi_user" gorm:"default:0"`   //is_multi_user==0>非单端口多用户承载端口  is_multi_user==1>混淆式单端口多用户承载端口  is_multi_user==2>协议式单端口多用户承载端口
}

// 用户与角色 多对多 表
type UserAndRole struct {
	UserID int
	RoleID int
}

// 用户校验邮箱
type UserRegisterEmail struct {
	UserName string `json:"user_name" binding:"required,email,max=40,min=8"` // 用户名
}

// 用户登录/重置密码 请求
type UserLogin struct {
	UserName  string `json:"user_name" binding:"required,email,max=40,min=8"` // 用户名
	Password  string `json:"password" binding:"required,max=20,min=4"`        // 密码
	EmailCode string `json:"email_code"`                                      //邮箱验证码
}

// 用户注册 请求
type UserRegister struct {
	UserName   string `json:"user_name" binding:"required,email,max=40,min=8"`              // 用户名
	Password   string `json:"password" binding:"required,max=20,min=4"`                     // 密码
	RePassword string `json:"re_password" binding:"required,eqfield=Password,max=20,min=4"` // 密码
	EmailCode  string `json:"email_code"`                                                   //邮箱验证码
}

// 新建用户/修改用户请求
type NewUser struct {
	User     User     `json:"user"      binding:"required"`
	RoleList []string `json:"check_list" binding:"required"` //选中的角色
}

// 修改密码 请求
type UserChangePassword struct {
	Password   string `json:"password" binding:"required,max=20,min=4"`                     // 密码
	RePassword string `json:"re_password" binding:"required,eqfield=Password,max=20,min=4"` // 密码
	EmailCode  string `json:"email_code"`
}

// users with total
type UsersWithTotal struct {
	Total    int64  `json:"total"`
	UserList []User `json:"user_list"`
}

// 可连接的用户 sspanel 响应模板
//
//	{
//		"ret": 1,
//		"data": [
//			{
//				"id": 1,
//				"email": "2493733085@qq.com",
//				"passwd": "Ws3cxaA5z5zZCLyZ",
//				"uuid": "9d280b53-ed44-38f5-8900-a6d35f3ba9b7",
//				"u": 2286096176,
//				"d": 28982317556,
//				"port": 1025,
//				"method": "rc4-md5",
//				"node_speedlimit": 0,
//				"node_connector": 3,
//				"protocol": "origin",
//				"protocol_param": "",
//				"obfs": "plain",
//				"obfs_param": "",
//				"forbidden_ip": "",
//				"forbidden_port": "",
//				"disconnect_ip": null,
//				"is_multi_user": 0
//			},
//		]
//	}
//
// 获取当前节点可连接的用户 响应
type SSUsers struct {
	ID             int    `json:"id"`
	UUID           string `json:"uuid"`
	U              int    `json:"u"`
	D              int    `json:"d"`
	NodeSpeedLimit int    `json:"node_speedlimit"`
	NodeConnector  int    `json:"node_connector"`
	// IsMultiUser    int    `json:"is_multi_user"` //用不到
}
