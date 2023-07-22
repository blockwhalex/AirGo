package model

import "time"

// TradeNo             string      `json:"trade_no"`              // 支付宝交易号
// OutTradeNo          string      `json:"out_trade_no"`          // 商户订单号
// BuyerLogonId        string      `json:"buyer_logon_id"`        // 买家支付宝账号
// TradeStatus         TradeStatus `json:"trade_status"`          // 交易状态
// TotalAmount         string      `json:"total_amount"`          // 订单金额
// ReceiptAmount       string      `json:"receipt_amount"`        // 实收金额
// BuyerPayAmount      string      `json:"buyer_pay_amount"`      // 付款金额
// Subject             string      `json:"subject"`               // 商品的标题/交易标题/订单标题/订单关键字等，是请求时对应的参数，原样通知回来。

type Orders struct {
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"-"        gorm:"index"`
	ID        int        `json:"id"       gorm:"primary_key"`
	UserID    int        `json:"user_id"   gorm:"comment:外键"` //外键
	UserName  string     `json:"user_name" gorm:"comment:订单拥有者"`
	User      User       `json:"user"     gorm:"foreignKey:UserID;references:ID"` //订单拥有者

	OutTradeNo string `json:"out_trade_no" gorm:"comment:商户订单号"` // 商户订单号
	GoodsID    int    `json:"goods_id"     gorm:"comment:商品id"`  //商品id
	Subject    string `json:"subject"      gorm:"comment:商品的标题"` // 商品的标题/交易标题/订单标题/订单关键字等，是请求时对应的参数，原样通知回来。
	Price      string `json:"price"        gorm:"comment:商品的价格"` //商品价格
	PayType    string `json:"pay_type"     gorm:"comment:支付方式"`  //支付方式 alipay wechat
	//Status     string `json:"status"       gorm:"default:created;comment:交易状态"`
	//支付宝参数
	QRCode         string `json:"qr_code"`                                                                          // 当前预下单请求生成的二维码码串，可以用二维码生成工具根据该码串值生成对应的二维码
	TradeNo        string `json:"trade_no"         gorm:"comment:支付宝交易号"`                                           // 支付宝交易号
	BuyerLogonId   string `json:"buyer_logon_id"   gorm:"comment:买家支付宝账号"`                                          // 买家支付宝账号
	TradeStatus    string `json:"trade_status"     gorm:"default:created;comment:交易状态"`                             // 交易状态 1、WAIT_BUYER_PAY（交易创建，等待买家付款）；2、TRADE_CLOSED（未付款交易超时关闭，或支付完成后全额退款）；3、TRADE_SUCCESS（交易支付成功）； 4、TRADE_FINISHED（交易结束，不可退款）；5、completed（手动完成订单）；6、created（订单已创建）
	TotalAmount    string `json:"total_amount"     gorm:"comment:订单金额"`                                             // 订单金额
	ReceiptAmount  string `json:"receipt_amount"   gorm:"comment:实收金额"`                                             // 实收金额
	BuyerPayAmount string `json:"buyer_pay_amount" gorm:"comment:付款金额"`                                             // 付款金额
	ProductCode    string `json:"product_code"     gorm:"default:FACE_TO_FACE_PAYMENT;comment:销售产品码，与支付宝签约的产品码名称。"` // 销售产品码，与支付宝签约的产品码名称。当面付alipay.trade.pay接口中，product_code为：FACE_TO_FACE_PAYMENT
}

// 分页订单数据，带总数
type OrdersWithTotal struct {
	OrderList []Orders `json:"order_list"`
	Total     int64    `json:"total"`
}

// 订单收入统计
type OrderStatistics struct {
	Total       int64   `json:"total"`
	TotalAmount float64 `json:"total_amount"`
}

type OrdersHeader struct {
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
	ID             string `json:"id"`
	UserID         string `json:"user_id"`
	UserName       string `json:"user_name"`
	OutTradeNo     string `json:"out_trade_no"`
	GoodsID        string `json:"goods_id"`
	Subject        string `json:"subject"`
	Price          string `json:"price"`
	PayType        string `json:"pay_type"`
	TradeNo        string `json:"trade_no"`
	BuyerLogonId   string `json:"buyer_logon_id"`
	TradeStatus    string `json:"trade_status"`
	TotalAmount    string `json:"total_amount"`
	ReceiptAmount  string `json:"receipt_amount"`
	BuyerPayAmount string `json:"buyer_pay_amount"`
}

var OrdersHeaderItem = OrdersHeader{
	CreatedAt:  "创建日期",
	UpdatedAt:  "更新日期",
	ID:         "ID",
	UserID:     "用户ID",
	UserName:   "用户名",
	OutTradeNo: "系统订单号",
	GoodsID:    "商品ID",
	Subject:    "商品标题",
	//Price:"价格",
	PayType:        "支付类型",
	TradeNo:        "支付宝订单号",
	BuyerLogonId:   "买家支付宝账号",
	TradeStatus:    "交易状态",
	TotalAmount:    "订单金额",
	ReceiptAmount:  "实收金额",
	BuyerPayAmount: "付款金额",
}
