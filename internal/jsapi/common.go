package jsapi

import "fmt"

const (
	CREATE_ORDER_URL = "https://api.mch.weixin.qq.com/v3/pay/transactions/jsapi"

	// example: https://api.mch.weixin.qq.com/v3/pay/transactions/id/1217752501201407033233368018?mchid=1230000109
	query_order_url = "https://api.mch.weixin.qq.com/v3/pay/transactions/id/%s?mchid=%s"

	// https://api.mch.weixin.qq.com/v3/pay/transactions/out-trade-no/{out_trade_no}/close
	CLOSE_ORDER_URL = "https://api.mch.weixin.qq.com/v3/pay/transactions/out-trade-no/%s/close"

	REFUND_URL = "https://api.mch.weixin.qq.com/v3/refund/domestic/refunds"

	// https://api.mch.weixin.qq.com/v3/refund/domestic/refunds/{out_refund_no}
	QUERY_REFUND_URL = "https://api.mch.weixin.qq.com/v3/refund/domestic/refunds/%s"
)

func GetQueryORderURL(transactionID, mchid string) string {
	return fmt.Sprintf(query_order_url, transactionID, mchid)
}

// todo 暂时不定义，不清楚错误结构
type ErrResp struct {

}