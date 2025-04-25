package user

import (
	"context"

	"TSAccountService/biz/bizcontext"
	user "TSAccountService/hertz_gen/user"

	"github.com/cloudwego/hertz/pkg/app"
)

type ChangePasswordService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewChangePasswordService(Context context.Context, RequestContext *app.RequestContext) *ChangePasswordService {
	return &ChangePasswordService{RequestContext: RequestContext, Context: Context}
}

func (h *ChangePasswordService) Run(ctx *bizcontext.BizContext, req *user.ChangePasswordReq) (resp *user.Empty, err error) {
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
