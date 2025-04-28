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

type UpdateUserInfoService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewUpdateUserInfoService(Context context.Context, RequestContext *app.RequestContext) *UpdateUserInfoService {
	return &UpdateUserInfoService{RequestContext: RequestContext, Context: Context}
}

func (h *UpdateUserInfoService) Run(ctx *bizcontext.BizContext, req *user.UpdateUserInfoReq) (resp *user.Empty, err error) {
	userBasic, err := model.FindUserByID(ctx.User.ID)
	if err != nil {
		hlog.CtxErrorf(ctx, "FindUserByID failed, err: %v", err)
		return nil, errno.UserNotFoundErr
	}

	if req.Nickname != "" {
		userBasic.Nickname = req.Nickname
	}

	if req.Avatar != "" {
		userBasic.Avatar = req.Avatar
	}

	err = model.UpdateUser(userBasic)
	if err != nil {
		hlog.CtxErrorf(ctx, "update user info failed, err: %v", err)
		return nil, errno.UpdateUserInfoErr
	}

	return &user.Empty{}, nil
}
