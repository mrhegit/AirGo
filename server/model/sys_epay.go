package model

// 易支付
type EpayOrder struct {
}

type EpayResponse struct {
	Pid         int64  //商户ID
	TradeNo     string //易支付订单号
	OutTradeNo  string //商户订单号
	Type        string //支付方式
	Name        string //商品名称
	Money       string //商品金额
	TradeStatus string //支付状态
	Param       string //业务扩展参数
	Sign        string //签名字符串
	SignType    string //签名类型
}
