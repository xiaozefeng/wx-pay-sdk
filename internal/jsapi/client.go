package jsapi

import (
	"context"
	"encoding/json"

	"github.com/wx-pay-sdk/internal/pkg/http"
)

// 创建预付单
func CreatePrepayOrder(ctx *context.Context, req *CreateOrderReq) (*CreateOrderResp, error) {
	b, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	result, err := http.PostJSON(ctx, CREATE_ORDER_URL, b)
	if err != nil {
		return nil, err
	}
	var resp CreateOrderResp
	err = json.Unmarshal(result, &resp)
	return &resp, err
}

// 查询订单
func QueryOrder(ctx *context.Context, )
