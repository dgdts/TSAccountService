package user

import (
	"context"

	"TSAccountService/biz/bizcontext"
	"TSAccountService/biz/errno"
	"TSAccountService/biz/model"
	user "TSAccountService/hertz_gen/user"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
)

type CheckPasswordSetService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCheckPasswordSetService(Context context.Context, RequestContext *app.RequestContext) *CheckPasswordSetService {
	return &CheckPasswordSetService{RequestContext: RequestContext, Context: Context}
}

func (h *CheckPasswordSetService) Run(ctx *bizcontext.BizContext, req *user.Empty) (resp *user.CheckPasswordSetResp, err error) {
	userBasic, err := model.FindUserByID(ctx.User.ID)
	if err != nil {
		hlog.CtxErrorf(ctx, "FindUserByID failed, err: %v", err)
		return nil, errno.UserNotFoundErr
	}

	isSet := userBasic.Password != ""
	resp = &user.CheckPasswordSetResp{
		IsSet: isSet,
	}

	return resp, nil
}
