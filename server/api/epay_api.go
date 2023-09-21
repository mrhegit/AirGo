package api

import (
	"AirGo/model"
	"AirGo/service"
	"AirGo/utils/other_plugin"
	"AirGo/utils/response"
	"fmt"
	"github.com/gin-gonic/gin"
)

// 易支付跳转通知

// 易支付异步通知
func EpayNotify(ctx *gin.Context) {
	uIDInt, _ := other_plugin.GetUserIDFromGinContext(ctx)

	var epayRes model.EpayResponse
	err := ctx.ShouldBind(&epayRes)
	if err != nil {
		return
	}
	fmt.Println("易支付异步通知", epayRes)
	if epayRes.TradeStatus == "TRADE_SUCCESS" {

		//处理订单，用户订阅
		service.UpdateOrderForEpay(&epayRes, uIDInt)

		//返回success以表示服务器接收到了订单通知
		response.OK("success", nil, ctx)
	}
	return
}
