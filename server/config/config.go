package config

// 全局变量-config配置
type Server struct {
	JWT     JWT     `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Email   Email   `mapstructure:"email" json:"email" yaml:"email"`
	System  System  `mapstructure:"system" json:"system" yaml:"system"`
	Captcha Captcha `mapstructure:"captcha" json:"captcha" yaml:"captcha"`
	Mysql   Mysql   `mapstructure:"mysql" json:"mysql" yaml:"mysql"` // gorm
	Pay     Pay     `mapstructure:"pay" json:"pay" yaml:"pay"`       // 支付相关配置
	Sqlite  Sqlite  `mapstructure:"sqlite" json:"sqlite" yaml:"sqlite"`
	// Zap     Zap     `mapstructure:"zap" json:"zap" yaml:"zap"`
	// Redis   Redis   `mapstructure:"redis" json:"redis" yaml:"redis"`
	//Cors CORS `mapstructure:"cors" json:"cors" yaml:"cors"`	// 跨域配置
}

// mysql配置
type Mysql struct {
	Path         string `mapstructure:"path" json:"path" yaml:"path"`                               // 服务器地址:端口
	Port         string `mapstructure:"port" json:"port" yaml:"port"`                               //:端口
	Config       string `mapstructure:"config" json:"config" yaml:"config"`                         // 高级配置
	Dbname       string `mapstructure:"db-name" json:"db-name" yaml:"db-name"`                      // 数据库名
	Username     string `mapstructure:"username" json:"username" yaml:"username"`                   // 数据库用户名
	Password     string `mapstructure:"password" json:"password" yaml:"password"`                   // 数据库密码
	Prefix       string `mapstructure:"prefix" json:"prefix" yaml:"prefix"`                         //全局表前缀，单独定义TableName则不生效
	Singular     bool   `mapstructure:"singular" json:"singular" yaml:"singular"`                   //是否开启全局禁用复数，true表示开启
	Engine       string `mapstructure:"engine" json:"engine" yaml:"engine" default:"InnoDB"`        //数据库引擎，默认InnoDB
	MaxIdleConns int    `mapstructure:"max-idle-conns" json:"max-idle-conns" yaml:"max-idle-conns"` // 空闲中的最大连接数
	MaxOpenConns int    `mapstructure:"max-open-conns" json:"max-open-conns" yaml:"max-open-conns"` // 打开到数据库的最大连接数
	LogMode      string `mapstructure:"log-mode" json:"log-mode" yaml:"log-mode"`                   // 是否开启Gorm全局日志
	LogZap       bool   `mapstructure:"log-zap" json:"log-zap" yaml:"log-zap"`                      // 是否通过zap写入日志文件
}
type Captcha struct {
	KeyLong            int `mapstructure:"key-long" json:"key-long" yaml:"key-long"`                                     // 验证码长度
	ImgWidth           int `mapstructure:"img-width" json:"img-width" yaml:"img-width"`                                  // 验证码宽度
	ImgHeight          int `mapstructure:"img-height" json:"img-height" yaml:"img-height"`                               // 验证码高度
	OpenCaptcha        int `mapstructure:"open-captcha" json:"open-captcha" yaml:"open-captcha"`                         // 防爆破验证码开启此数，0代表每次登录都需要验证码，其他数字代表错误密码此数，如3代表错误三次后出现验证码
	OpenCaptchaTimeOut int `mapstructure:"open-captcha-timeout" json:"open-captcha-timeout" yaml:"open-captcha-timeout"` // 防爆破验证码超时时间，单位：s(秒)
}
type JWT struct {
	SigningKey  string `mapstructure:"signing-key" json:"signing-key" yaml:"signing-key"`    // jwt签名
	ExpiresTime string `mapstructure:"expires-time" json:"expires-time" yaml:"expires-time"` // 过期时间
	BufferTime  string `mapstructure:"buffer-time" json:"buffer-time" yaml:"buffer-time"`    // 缓冲时间
	Issuer      string `mapstructure:"issuer" json:"issuer" yaml:"issuer"`                   // 签发者
}
type System struct {
	Port          int    `mapstructure:"port" json:"addr" yaml:"addr"`                               // 端口值
	DbType        string `mapstructure:"db-type" json:"db-type" yaml:"db-type"`                      // 数据库类型:mysql(默认)|sqlite|sqlserver|postgresql
	UseMultipoint bool   `mapstructure:"use-multipoint" json:"use-multipoint" yaml:"use-multipoint"` // 多点登录拦截
	SubName       string `mapstructure:"sub-name" json:"sub-name" yaml:"sub-name"`
}
type Email struct {
	From     string `mapstructure:"from" json:"from" yaml:"from"`             // 发件人
	Secret   string `mapstructure:"secret" json:"secret" yaml:"secret"`       // 密钥
	Host     string `mapstructure:"host" json:"host" yaml:"host"`             // 服务器地址
	Port     int    `mapstructure:"port" json:"port" yaml:"port"`             // 端口
	IsSSL    bool   `mapstructure:"is-ssl" json:"is-ssl" yaml:"is-ssl"`       // 是否SSL
	Nickname string `mapstructure:"nickname" json:"nickname" yaml:"nickname"` // 昵称
}

// 支付相关配置
type Pay struct {
	ReturnURL    string `mapstructure:"returnURL" yaml:"returnURL"`
	AppID        string `mapstructure:"appID" yaml:"appID"`
	PrivateKey   string `mapstructure:"privateKey" yaml:"privateKey"`
	AliPublicKey string `mapstructure:"aliPublicKey" yaml:"aliPublicKey"`
	EncryptKey   string `mapstructure:"encryptKey" yaml:"encryptKey"`
}
type Sqlite struct {
	Path string `mapstructure:"path" yaml:"path"`
}
