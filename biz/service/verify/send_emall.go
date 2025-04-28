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

type SendEmallService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewSendEmallService(Context context.Context, RequestContext *app.RequestContext) *SendEmallService {
	return &SendEmallService{RequestContext: RequestContext, Context: Context}
}

func (h *SendEmallService) Run(ctx *bizcontext.BizContext, req *verify.SendEmallReq) (resp *verify.Empty, err error) {
	var code string
	if biz_config.GetBizConfig().MockConfig.Enable {
		code = biz_config.GetBizConfig().MockConfig.EmailCode
	} else {
		code, err = random.GenerateVerifyCode()
		if err != nil {
			return nil, err
		}
	}

	err = verify_code.StoreVerifyCode(ctx, req.Emall, code, h.RequestContext.ClientIP())

	if err != nil {
		return nil, err
	}

	// TODO: send email
	hlog.CtxDebugf(ctx, "send email email:%s, code:%s", req.Emall, code)

	return &verify.Empty{}, nil
}
