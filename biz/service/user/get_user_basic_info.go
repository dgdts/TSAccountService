package user

import (
	"context"

	"TSAccountService/biz/bizcontext"
	user "TSAccountService/hertz_gen/user"

	"github.com/cloudwego/hertz/pkg/app"
)

type GetUserBasicInfoService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetUserBasicInfoService(Context context.Context, RequestContext *app.RequestContext) *GetUserBasicInfoService {
	return &GetUserBasicInfoService{RequestContext: RequestContext, Context: Context}
}

func (h *GetUserBasicInfoService) Run(ctx *bizcontext.BizContext, req *user.Empty) (resp *user.GetUserBasicInfoResp, err error) {
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
