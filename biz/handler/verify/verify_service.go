package verify

import (
	"context"

	"TSAccountService/biz/response"
	logic "TSAccountService/biz/service/verify"
	_ "TSAccountService/hertz_gen/verify"

	"github.com/cloudwego/hertz/pkg/app"
)

// SendSMS .
// @router /api/v1/verify/send-sms [POST]
func SendSMS(ctx context.Context, c *app.RequestContext) {
	response.JSON(ctx, c, logic.NewSendSMSService(ctx, c).Run)
}

// SendEmall .
// @router /api/v1/verify/send-email [POST]
func SendEmall(ctx context.Context, c *app.RequestContext) {
	response.JSON(ctx, c, logic.NewSendEmallService(ctx, c).Run)
}
