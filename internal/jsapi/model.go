package jsapi

type CreateOrderReq struct {

	// {
	// 	"mchid": "1900006XXX",
	// 	"out_trade_no": "1217752501201407033233368318",
	// 	"appid": "wxdace645e0bc2cXXX",
	// 	"description": "Image形象店-深圳腾大-QQ公仔",
	// 	"notify_url": "https://weixin.qq.com/",
	// 	"amount": {
	// 		"total": 1,
	// 		"currency": "CNY"
	// 	},
	// 	"payer": {
	// 		"openid": "o4GgauInH_RCEdvrrNGrntXDuXXX"
	// 	}
	// }
	AppID       string `json:"appid,omitempty"`
	MchId       string `json:"mchid,omitempty"`
	OutTradeNo  string `json:"out_trade_no,omitempty"`
	Description string `json:"description,omitempty"`
	NotifyURL   string `json:"notify_url,omitempty"`
	Amount      struct {
		Total    int    `json:"total,omitempty"`
		Currency string `json:"currency,omitempty"`
	} `json:"amount,omitempty"`
	Payer      payer      `json:"payer,omitempty"`
	Attach     string     `json:"attach,omitempty"`
	GoodsTag   string     `json:"goods_tag,omitempty"`
	SettleInfo settleinfo `json:"settle_info,omitempty"`
}

// 结算信息
type settleinfo struct {
	ProfitSharing bool `json:"profit_sharing,omitempty"`
}

// 支付人信息
type payer struct {
	OpenID string `json:"openid,omitempty"`
}

type CreateOrderResp struct {
	PrepayID string `json:"prepay_id,omitempty"`
}

// ----------------

// 优先使用 TransactionID, 如果TransactionID为空，才使用 OutOrderNo
type QueryOrderReq struct {
	// 微信支付系统生成的订单号
	TransactionID string `json:"transaction_id,omitempty"`
	// 商家订单号
	OutOrderNo string `json:"out_order_no,omitempty"`
	MchID      string `json:"mch_id,omitempty"`
}

type QueryORderResp struct {
	Amount struct {
		Currency      string `json:"currency,omitempty"`
		PayerCurrency string `json:"payer_currency,omitempty"`
		PayerTotal    int    `json:"payer_total,omitempty"`
		Total         int    `json:"total,omitempty"`
	} `json:"amount,omitempty"`
	Appid          string `json:"appid,omitempty"`
	Attach         string `json:"attach,omitempty"`
	BankType       string `json:"bank_type,omitempty"`
	Mchid          string `json:"mchid,omitempty"`
	OutTradeNo     string `json:"out_trade_no,omitempty"`
	Payer          payer  `json:"payer,omitempty"`
	SuccessTime    string `json:"success_time,omitempty"`
	TradeState     string `json:"trade_state,omitempty"`
	TradeStateDesc string `json:"trade_state_desc,omitempty"`
	TradeType      string `json:"trade_type,omitempty"`
	TransactionID  string `json:"transaction_id,omitempty"`
}
