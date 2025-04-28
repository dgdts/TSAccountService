package user

import (
	"context"

	"TSAccountService/biz/bizcontext"
	"TSAccountService/biz/config/biz_config"
	"TSAccountService/biz/constant"
	"TSAccountService/biz/errno"
	user "TSAccountService/hertz_gen/user"

	"TSAccountService/biz/utils"

	"github.com/cloudwego/hertz/pkg/app"
)

type LogoutService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewLogoutService(Context context.Context, RequestContext *app.RequestContext) *LogoutService {
	return &LogoutService{RequestContext: RequestContext, Context: Context}
}

func (h *LogoutService) Run(ctx *bizcontext.BizContext, req *user.Empty) (resp *user.Empty, err error) {
	token := h.RequestContext.Request.Header.Get(constant.HeaderToken)
	if token == "" {
		return nil, errno.TokenEmptyErr
	}
	err = utils.SetLogout(ctx, token, biz_config.GetBizConfig().JWTConfig.ExpireMinute)

	return &user.Empty{}, err
}
