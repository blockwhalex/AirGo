package alipay_plugin

import (
	"AirGo/global"
	"AirGo/model"
	"fmt"
	"strconv"

	"github.com/smartwalle/alipay/v3"
)

func InitAlipayClient() (*alipay.Client, error) {

	//false时用开发网关，https://openapi.alipaydev.com/gateway.do；true时用正式环境网关，https://openapi.alipay.com/gateway.do
	client, err := alipay.New(global.Server.Pay.AppID, global.Server.Pay.PrivateKey, true)
	if err != nil {
		fmt.Println("初始化支付宝失败, 错误信息为", err)
		//os.Exit(-1)
		return nil, err
	}

	// 加载内容密钥（可选），详情查看 https://opendocs.alipay.com/common/02mse3
	client.SetEncryptKey(global.Server.Pay.EncryptKey)

	// 下面两种方式只能二选一
	var cert = false
	if cert {
		// 使用支付宝证书
		fmt.Println("加载证书", client.LoadAppPublicCertFromFile("appPublicCert.crt"))
		fmt.Println("加载证书", client.LoadAliPayRootCertFromFile("alipayRootCert.crt"))
		fmt.Println("加载证书", client.LoadAliPayPublicCertFromFile("alipayPublicCert.crt"))
	} else {
		// 使用支付宝公钥
		fmt.Println("加载公钥", client.LoadAliPayPublicKey(global.Server.Pay.AliPublicKey))
	}
	return client, nil
}

// alipay.trade.precreate(统一收单线下交易预创建),生成二维码后，展示给用户，由用户扫描二维码完成订单支付（当面付）
func TradePreCreatePay(client *alipay.Client, sysOrder *model.Orders) (*alipay.TradePreCreateRsp, error) {
	//创建支付宝订单 请求模板
	// order := alipay.TradePreCreate{
	// 	Trade: alipay.Trade{
	// 		//NotifyURL:  "" ,  //异步通知的URL，该URL是支付宝服务器端自动调用商户服务端接口的地址，支付成功后调用，再根据支付宝转发的参数进行订单状态处理,异步的post请求
	// 		//ReturnURL: "http://", //同步返回URL，是一个页面跳转通知（支付成功后，从支付宝跳转到指定的地址）。同步的get请求
	// 		//AppAuthToken: "", // 可选

	// 		// biz content，这四个参数是必须的
	// 		Subject:     "一个馒头",                 // 订单标题
	// 		OutTradeNo:  "1008610010",           // 商户订单号，64个字符以内、可包含字母、数字、下划线；需保证在商户端不重复
	// 		TotalAmount: "0.01",                 // 订单总金额，单位为元，精确到小数点后两位，取值范围[0.01,100000000]
	// 		ProductCode: "FACE_TO_FACE_PAYMENT", // 销售产品码，与支付宝签约的产品码名称。 参考官方文档,
	// 		//App 支付时默认值为 QUICK_MSECURITY_PAY
	// 		//手机网站支付产品alipay.trade.wap.pay接口中，product_code为：QUICK_WAP_WAY
	// 		//电脑网站支付产品alipay.trade.page.pay接口中，product_code为：FAST_INSTANT_TRADE_PAY
	// 		//当面付条码支付产品alipay.trade.pay接口中，product_code为：FACE_TO_FACE_PAYMENT
	// 	},
	// }
	//创建支付宝订单
	var order alipay.TradePreCreate
	//order.NotifyURL = global.Server.Pay.ReturnURL  //支付结果放在轮询里判断
	goodsID := strconv.Itoa(sysOrder.GoodsID)
	order.Subject = goodsID + "-" + sysOrder.Subject
	order.OutTradeNo = sysOrder.OutTradeNo
	order.TotalAmount = sysOrder.TotalAmount
	order.ProductCode = sysOrder.ProductCode
	res, err := client.TradePreCreate(order)
	return res, err
	//响应模板
	// 	{
	// 	"code": 0,
	// 	"msg": "alipay TradePreCreatePay success:",
	// 	"data": {
	// 		"alipay_trade_precreate_response": {
	// 			"code": "10000",
	// 			"msg": "Success",
	// 			"sub_code": "",
	// 			"sub_msg": "",
	// 			"out_trade_no": "5",
	// 			"qr_code": "https://qr.alipay.com/bax07220zdz0k58x5abw2504"
	// 		},
	// 		"sign": "EmZmz7Jix2GLtScaysE9D0DF9Sw9ZDuuums7CXywFO83g/dnOasZiAQnDhsgoMq9JmPnygIog4+myEcxXqmoLM2qZX2zy3Aof7CbVzLwA931kq09u6y54h28R+BvILLZAR5gmSYW2oh4/gWO24yK8awHLndCAQhNuHFOkMwCAcDRKGjhKkDb9XIx/do99V/Xa9w8pJhHSt1ONaIjyWufK6b4YcVg3bGldBTG+xpqDvzXSYFc5lBRfgAJxn8NklTKVj/PLFr3nM4IJ/fYFaJuHS2/pjQThyDiEsjPvEhA9aPEeXK03J8Qea0HFAuM9i2kw1OqmeN0oiHCrVVSCFGPRg=="
	// 	}
	// }
	//响应模板
	// type TradePreCreateRsp struct {
	// 	Content struct {
	// 		Code       Code   `json:"code"`
	// 		Msg        string `json:"msg"`
	// 		SubCode    string `json:"sub_code"`
	// 		SubMsg     string `json:"sub_msg"`
	// 		OutTradeNo string `json:"out_trade_no"` // 创建交易传入的商户订单号
	// 		QRCode     string `json:"qr_code"`      // 当前预下单请求生成的二维码码串，可以用二维码生成工具根据该码串值生成对应的二维码
	// 	} `json:"alipay_trade_precreate_response"`
	// 	Sign string `json:"sign"`
	// }

	// if err != nil {
	// 	fmt.Println(err)
	// }
	//fmt.Println("统一收单交易支付接口TradePay返回:", res)

}

//func TradeCreate(client *alipay.Client, sysOrder *model.Orders) {
//	client.TradeCreate()
//}

// 统一收单线下交易查询
func TradeQuery(client *alipay.Client, sysOrder *model.Orders) (*alipay.TradeQueryRsp, error) {
	// TradeQuery 统一收单线下交易查询接口请求参数 https://docs.open.alipay.com/api_1/alipay.trade.query/
	//type TradeQuery struct {
	//	AppAuthToken string   `json:"-"`                       // 可选
	//	OutTradeNo   string   `json:"out_trade_no,omitempty"`  // 订单支付时传入的商户订单号, 与 TradeNo 二选一
	//	TradeNo      string   `json:"trade_no,omitempty"`      // 支付宝交易号
	//	QueryOptions []string `json:"query_options,omitempty"` // 可选 查询选项，商户通过上送该字段来定制查询返回信息 TRADE_SETTLE_INFO(交易结算信息)
	//}
	var p = alipay.TradeQuery{
		OutTradeNo: sysOrder.OutTradeNo,
		//QueryOptions: []string{"TRADE_SETTLE_INFO"},
	}
	//fmt.Println("统一收单线下交易查询 p:", p)
	rsp, err := client.TradeQuery(p)
	return rsp, err

	//  if rsp.Content.Code != alipay.CodeSuccess
	//	fmt.Println(rsp.Content.Code, rsp.Content.Msg, rsp.Content.SubMsg)
	//Code   Msg                SubCode                    SubMsg       TradeNo                       OutTradeNo                BuyerLogonId   TradeStatus
	//40004 Business  Failed    ACQ.TRADE_NOT_EXIST        交易不存在                                  168425758005579100000000
	//10000 Success                                                     2023051822001475841447588320  168438189998030100010000 249***@qq.com   WAIT_BUYER_PAY
	//10000 Success                                                     2023051822001475841447588320  168438189998030100010000 249***@qq.com   TRADE_SUCCESS

}

// 统一收单交易关闭接口
func TradeClose(client *alipay.Client, sysOrder *model.Orders) (*alipay.TradeCloseRsp, error) {
	//type TradeClose struct {
	//	AppAuthToken string `json:"-"`                      // 可选
	//	NotifyURL    string `json:"-"`                      // 可选
	//	OutTradeNo   string `json:"out_trade_no,omitempty"` // 与 TradeNo 二选一
	//	TradeNo      string `json:"trade_no,omitempty"`     // 与 OutTradeNo 二选一
	//	OperatorId   string `json:"operator_id,omitempty"`  // 可选
	//}
	var p = alipay.TradeClose{
		OutTradeNo: sysOrder.OutTradeNo,
	}
	rsp, err := client.TradeClose(p)
	return rsp, err
}
