package verify

import (
	"context"

	"TSAccountService/biz/bizcontext"
	"TSAccountService/biz/config/biz_config"
	verify "TSAccountService/hertz_gen/verify"
	"TSAccountService/kit/random"
	"TSAccountService/kit/verify_code"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
)

type SendSMSService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewSendSMSService(Context context.Context, RequestContext *app.RequestContext) *SendSMSService {
	return &SendSMSService{RequestContext: RequestContext, Context: Context}
}

func (h *SendSMSService) Run(ctx *bizcontext.BizContext, req *verify.SendSMSReq) (resp *verify.Empty, err error) {
	var code string
	if biz_config.GetBizConfig().MockConfig.Enable {
		code = biz_config.GetBizConfig().MockConfig.SMSCode
	} else {
		code, err = random.GenerateVerifyCode()
		if err != nil {
			return nil, err
		}
	}

	err = verify_code.StoreVerifyCode(ctx, req.Phone, code, h.RequestContext.ClientIP())

	if err != nil {
		return nil, err
	}

	// TODO: send sms
	hlog.CtxDebugf(ctx, "send sms phone:%s, code:%s", req.Phone, code)

	return &verify.Empty{}, nil
}
