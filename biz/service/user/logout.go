package user

import (
	"context"

	"TSAccountService/biz/bizcontext"
	user "TSAccountService/hertz_gen/user"

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
