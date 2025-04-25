package user

import (
	"context"

	"TSAccountService/biz/bizcontext"
	user "TSAccountService/hertz_gen/user"

	"github.com/cloudwego/hertz/pkg/app"
)

type RegisterWithThirdPartyService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewRegisterWithThirdPartyService(Context context.Context, RequestContext *app.RequestContext) *RegisterWithThirdPartyService {
	return &RegisterWithThirdPartyService{RequestContext: RequestContext, Context: Context}
}

func (h *RegisterWithThirdPartyService) Run(ctx *bizcontext.BizContext, req *user.RegisterWithThirdPartyReq) (resp *user.Empty, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	// define your error in errno
	// if err != nil {
	// 	return nil, err
	// }
	return
}
