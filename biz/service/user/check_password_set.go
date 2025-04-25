package user

import (
	"context"

	"TSAccountService/biz/bizcontext"
	user "TSAccountService/hertz_gen/user"

	"github.com/cloudwego/hertz/pkg/app"
)

type CheckPasswordSetService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCheckPasswordSetService(Context context.Context, RequestContext *app.RequestContext) *CheckPasswordSetService {
	return &CheckPasswordSetService{RequestContext: RequestContext, Context: Context}
}

func (h *CheckPasswordSetService) Run(ctx *bizcontext.BizContext, req *user.Empty) (resp *user.CheckPasswordSetResp, err error) {
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
