package model

import "time"

// sspanel 响应 获取当前请求节点的节点设置
type SSNodeInfo struct {
	// NodeGroup      int    `json:"node_group"`
	// NodeClass      int    `json:"node_class"`
	//MuOnly         int    `json:"mu_only"` //只启用单端口多用户
	NodeSpeedlimit int    `json:"node_speedlimit"` //节点限速/Mbps
	TrafficRate    int    `json:"traffic_rate"`    //倍率
	Sort           int    `json:"sort"`            //类型
	Server         string `json:"server"`          //
	SSType         string `json:"type"`            //显示与隐藏
}
type NodeOld struct {
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"-"            gorm:"index"`
	ID        int        `json:"id"       gorm:"primary_key"`
	//节点基础参数
	// NodeGroup int `json:"node_group"`
	// NodeClass int `json:"node_class"`
	NodeOrder       int    `json:"node_order"`                           //节点排序
	Sort            int    `json:"sort"            gorm:"default:11"`    //类型sort==11  V2Ray vmess
	Name            string `json:"name"`                                 //节点名
	Address         string `json:"address"`                              //节点地址
	Port            string `json:"port"`                                 //端口
	NodeSpeedlimit  int    `json:"node_speedlimit" gorm:"default:0"`     //限速
	TrafficRate     int    `json:"traffic_rate"    gorm:"default:1"`     //倍率
	Status          bool   `json:"status"          gorm:"default:true"`  //节点是否启用，true启用，flase禁用
	EnableTransfer  bool   `json:"enable_transfer" gorm:"default:false"` //是否启用中转
	TransferAddress string `json:"transfer_address"`                     //中转ip
	TransferPort    string `json:"transfer_port"`                        //中转port
	TotalUp         int    `json:"total_up"        gorm:"-"`             //
	TotalDown       int    `json:"total_down"      gorm:"-"`             //
	//vmess参数
	V            string `json:"v"        gorm:"default:2"`                //
	Aid          string `json:"aid"      gorm:"default:0"`                //额外ID
	Scy          string `json:"scy"      gorm:"default:auto"`             //加密方式 （auto none chacha20-poly1305 aes-128-gcm zero）    vless字段：encryption
	Net          string `json:"net"      gorm:"default:ws"`               //传输协议 （ws tcp kcp quic grpc）                            vless字段：type
	Disguisetype string `json:"type"     gorm:"default:none"`             //伪装类型 （none http）                                       vless字段：headerType
	Host         string `json:"host"     gorm:"default:tms.dingtalk.com"` //伪装域名                                                     vless字段：host
	Path         string `json:"path"     gorm:"default:/"`                //path                                                       vless字段：path
	Tls          string `json:"tls"`                                      //传输层安全 （为空 tls）                                        vless字段：security
	Sni          string `json:"sni"`                                      //默认为空                                                     vless字段：sni
	Alpn         string `json:"alpn"`
	Fp           string `json:"fp"`
	//其他参数
	Goods       []Goods      `json:"goods"         gorm:"many2many:goods_and_nodes"` //多对多,关联商品
	TrafficLogs []TrafficLog `json:"-"   gorm:"foreignKey:NodeID;references:ID"`     //has many
}

// sspanel vmess 格式
type Vmess struct {
	V            string `json:"v" `   //
	Name         string `json:"ps"`   //节点名
	Address      string `json:"add"`  //节点地址
	Port         string `json:"port"` //端口
	Uuid         string `json:"id"`   //用户UUID
	Aid          string `json:"aid"`  //额外ID
	Net          string `json:"net"`  //传输协议
	Disguisetype string `json:"type"` //伪装类型
	Host         string `json:"host"` //伪装域名
	Path         string `json:"path"` //
	Tls          string `json:"tls"`  //传输层安全
}

// clash  yaml格式
type ClashYaml struct {
	Port               int               `yaml:"port"`
	SocksPort          int               `yaml:"socks-port"`
	RedirPort          int               `yaml:"redir-port"`
	AllowLan           bool              `yaml:"allow-lan"`
	Mode               string            `yaml:"mode"`
	LogLevel           string            `yaml:"log-level"`
	ExternalController string            `yaml:"external-controller"`
	Secret             string            `yaml:"secret"`
	Proxies            []ClashProxy      `yaml:"proxies"`
	ProxyGroups        []ClashProxyGroup `yaml:"proxy-groups"`
	Rules              []string          `yaml:"rules"`
}
type ClashProxy struct {
	Name      string    `yaml:"name"`
	Type      string    `yaml:"type"`
	Server    string    `yaml:"server"`
	Port      string    `yaml:"port"`
	Uuid      string    `yaml:"uuid"`
	Alterid   string    `yaml:"alterId"`
	Cipher    string    `yaml:"cipher"`
	Udp       bool      `yaml:"udp"`
	Network   string    `yaml:"network"`
	WsPath    string    `yaml:"ws-path"`
	WsHeaders WsHeaders `yaml:"ws-headers"`
	WsOpts    WsOpts    `yaml:"ws-opts"`
	Tls       bool      `yaml:"tls"`
	Sni       string    `yaml:"sni"`
	Password  string    `yaml:"password"` //trojan 参数

}

type WsHeaders struct {
	Host string `yaml:"Host"`
}
type WsOpts struct {
	Path    string            `yaml:"path"`
	Headers map[string]string `yaml:"headers"`
}

type ClashProxyGroup struct {
	Name    string   `yaml:"name"`
	Type    string   `yaml:"type"`
	Proxies []string `yaml:"proxies"`
}

// 修改混淆
type SubHost struct {
	Host string `json:"host"`
}

// 查询节点 with total
type NodesWithTotal struct {
	NodeList []Node `json:"node_list"`
	Total    int64  `json:"total"`
}

// 节点状态
type NodeStatus struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Status      bool      `json:"status"`
	LastTime    time.Time `json:"last_time"`
	UserAmount  int       `json:"user_amount"`
	TrafficRate int       `json:"traffic_rate"`
	U           float64   `json:"u"`
	D           float64   `json:"d"`
}
