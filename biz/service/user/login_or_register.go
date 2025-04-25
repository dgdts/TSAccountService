package user

import (
	"context"

	"TSAccountService/biz/bizcontext"
	"TSAccountService/biz/utils/third_party_auth"
	user "TSAccountService/hertz_gen/user"

	"github.com/cloudwego/hertz/pkg/app"
)

type LoginOrRegisterService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewLoginOrRegisterService(Context context.Context, RequestContext *app.RequestContext) *LoginOrRegisterService {
	return &LoginOrRegisterService{RequestContext: RequestContext, Context: Context}
}

func (h *LoginOrRegisterService) Run(ctx *bizcontext.BizContext, req *user.LoginOrRegisterReq) (resp *user.LoginOrRegisterResp, err error) {
	if req.AuthType != "" {
		return nil, third_party_auth.Oauth(ctx, req.AuthType)
	}

	if req.Phone != "" {
		return nil, nil
	}

	return nil, nil
}
