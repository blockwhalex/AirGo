package model

import "time"

type Node struct {
	ID        int       `json:"id" gorm:"primarykey"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	//DeletedAt gorm.DeletedAt `json:"-"  gorm:"index"`

	//sspanel 参数
	// NodeGroup      int    `json:"node_group"`
	// NodeClass      int    `json:"node_class"`
	//MuOnly         int    `json:"mu_only"` //只启用单端口多用户
	NodeSpeedlimit int    `json:"node_speedlimit"` //节点限速/Mbps
	TrafficRate    int    `json:"traffic_rate"`    //倍率
	Sort           int    `json:"sort"`            //类型 vless(15) vmess(11) trojan(14)
	Server         string `json:"server"`          //
	//SSType         string `json:"type"`            //显示与隐藏

	//基础参数
	Remarks string `json:"remarks"` //别名
	Address string `json:"address"` //地址
	Port    int    `json:"port"`    //端口
	//Ns         string  `json:"ns"`         //ip地址
	//TcpingData float64 `json:"tcping"`     //延迟测试
	//Ascription string  `json:"ascription"` //abroad domestic
	NodeOrder int  `json:"node_order"` //节点排序
	Enabled   bool `json:"enabled"`    //是否为激活节点
	//中转参数
	EnableTransfer  bool   `json:"enable_transfer" gorm:"default:false"` //是否启用中转
	TransferAddress string `json:"transfer_address"`                     //中转ip
	TransferPort    int    `json:"transfer_port"`                        //中转port
	//上行/下行
	TotalUp   int `json:"total_up"        gorm:"-"` //
	TotalDown int `json:"total_down"      gorm:"-"` //
	//关联参数
	Goods       []Goods      `json:"goods"         gorm:"many2many:goods_and_nodes"` //多对多,关联商品
	TrafficLogs []TrafficLog `json:"-"   gorm:"foreignKey:NodeID;references:ID"`     //has many

	//vmess参数
	V   string `json:"v"   gorm:"default:2"`
	Scy string `json:"scy" gorm:"default:auto"` //加密方式 auto,none,chacha20-poly1305,aes-128-gcm,zero
	Aid int    `json:"aid" gorm:"default:0"`    //额外ID
	//vless参数
	VlessFlow       string `json:"flow"`       //流控 none,xtls-rprx-vision,xtls-rprx-vision-udp443
	VlessEncryption string `json:"encryption"` //加密方式 none

	//传输参数
	Network  string `json:"network" gorm:"default:ws"`    //传输协议 tcp,kcp,ws,h2,quic,grpc
	Type     string `json:"type"    gorm:"default:none"`  //伪装类型 ws,h2：无    tcp,kcp：none，http    quic：none，srtp，utp，wechat-video，dtls，wireguard
	Host     string `json:"host"`                         //伪装域名
	Path     string `json:"path"    gorm:"default:/"`     //path
	GrpcMode string `json:"mode"    gorm:"default:multi"` //grpc传输模式 gun，multi

	//传输层安全
	Security      string `json:"security"`                          //传输层安全类型 none,tls,reality
	Sni           string `json:"sni"`                               //
	Fingerprint   string `json:"fp"`                                //
	Alpn          string `json:"alpn"`                              //tls
	AllowInsecure bool   `json:"allowInsecure" gorm:"default:true"` //tls 跳过证书验证

	PublicKey string `json:"pbk"` //reality
	ShortId   string `json:"sid"` //reality
	SpiderX   string `json:"spx"` //reality
}
