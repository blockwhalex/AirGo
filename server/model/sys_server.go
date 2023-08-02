package model

import "time"

// Server 全局配置
type Server struct {
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"-" gorm:"index"`
	ID        int        `json:"id"   gorm:"primary_key"`

	JWT             JWT             `json:"jwt"      gorm:"embedded"`
	System          System          `json:"system"   gorm:"embedded"`
	Captcha         Captcha         `json:"captcha"  gorm:"embedded"` //验证码
	Pay             Pay             `json:"pay"      gorm:"embedded"` // 支付相关配置
	Email           Email           `json:"email"    gorm:"embedded"`
	RateLimitParams RateLimitParams `json:"rate_limit_params"  gorm:"embedded"`
	//Mysql   Mysql   `json:"mysql"    gorm:"embedded"` // gorm
}

type Email struct {
	EmailFrom     string `json:"email_from"`                                                          // 发件人
	EmailSecret   string `json:"email_secret"`                                                        // 密钥
	EmailHost     string `json:"email_host"`                                                          // 服务器地址
	EmailPort     int    `json:"email_port"`                                                          // 端口
	EmailIsSSL    bool   `json:"email_is_ssl"`                                                        // 是否SSL
	EmailNickname string `json:"email_nickname"`                                                      // 昵称
	EmailSubject  string `json:"email_subject" gorm:"default:hello!"`                                 //邮件主题
	EmailContent  string `json:"email_content" gorm:"default:<h1>验证码</h1><p>emailcode</p>;size:5000"` //邮件内容
}

type Captcha struct {
	KeyLong            int `json:"key_long"        gorm:"default:6"`         // 验证码长度
	ImgWidth           int `json:"img_width"       gorm:"default:240"`       // 验证码宽度
	ImgHeight          int `json:"img_height"      gorm:"default:80"`        // 验证码高度
	OpenCaptcha        int `json:"open_captcha"    gorm:"default:2"`         // 防爆破验证码开启此数，0代表每次登录都需要验证码，其他数字代表错误密码此数，如3代表错误三次后出现验证码
	OpenCaptchaTimeOut int `json:"open_captcha_time_out" gorm:"default:300"` // 防爆破验证码超时时间，单位：s(秒)
}
type JWT struct {
	SigningKey  string `json:"signing_key"  gorm:"default:AirGo"` // jwt签名
	ExpiresTime string `json:"expires_time" gorm:"default:7d"`    // 过期时间
	BufferTime  string `json:"buffer_time"  gorm:"default:1d"`    // 缓冲时间
	Issuer      string `json:"issuer"       gorm:"default:AirGo"` // 签发者
}
type System struct {
	EnableRegister       bool    `json:"enable_register"         gorm:"default:true"`  // 是否开启注册
	EnableEmailCode      bool    `json:"enable_email_code"       gorm:"default:false"` // 是否开启注册email 验证码
	EnableLoginEmailCode bool    `json:"enable_login_email_code" gorm:"default:false"` // 是否开启登录email 验证码
	IsMultipoint         bool    `json:"is_multipoint"     gorm:"default:true"`        // 是否多点登录
	SubName              string  `json:"sub_name"          gorm:"default:AirGo"`       // 订阅名称
	SubUrlRre            string  `json:"sub_url_pre"`                                  //订阅前缀
	MuKey                string  `json:"muKey"             gorm:"comment:前后端通信密钥"`     // 前后端通信密钥
	DefaultGoods         string  `json:"default_goods"     gorm:"comment:新用户默认套餐"`
	EnabledRebate        bool    `json:"enabled_rebate"    gorm:"default:false;comment:是否开启返利"`
	RebateRate           float64 `json:"rebate_rate"       gorm:"default:0.1;comment:返利率"`
	EnabledDeduction     bool    `json:"enabled_deduction" gorm:"default:false;comment:是否开启旧套餐抵扣"`
	DeductionThreshold   float64 `json:"deduction_threshold" gorm:"default:0.8;comment:旧套餐抵扣阈值,大于该值则抵扣"`
}

// Pay 支付相关配置
type Pay struct {
	ReturnURL    string `json:"return_url"`
	AppID        string `json:"app_id"`
	PrivateKey   string `json:"private_key"    gorm:"size:3000"`
	AliPublicKey string `json:"ali_public_key" gorm:"size:3000"`
	EncryptKey   string `json:"encrypt_key"`
}

// RateLimitParams 限流参数
type RateLimitParams struct {
	IPRoleParam int `json:"ip_role_param" gorm:"default:600"`
	VisitParam  int `json:"visit_param"   gorm:"default:60"`
}

// 公共配置参数
type PublicSystem struct {
	RebateRate float64 `json:"rebate_rate"` //佣金率
	SubUrlRre  string  `json:"sub_url_pre"` //订阅前缀

}
