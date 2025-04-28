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

type ChangePasswordService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewChangePasswordService(Context context.Context, RequestContext *app.RequestContext) *ChangePasswordService {
	return &ChangePasswordService{RequestContext: RequestContext, Context: Context}
}

func (h *ChangePasswordService) Run(ctx *bizcontext.BizContext, req *user.ChangePasswordReq) (resp *user.Empty, err error) {
	userBasic, err := model.FindUserByID(ctx.User.ID)
	if err != nil {
		return nil, errno.UserNotFoundErr
	}

	if userBasic.Password != req.OldPassword {
		return nil, errno.OldPasswordErr
	}

	userBasic.Password = req.NewPassword
	err = model.UpdateUser(userBasic)
	if err != nil {
		hlog.CtxErrorf(ctx, "update user password failed, err: %v", err)
		return nil, errno.InternalErr
	}

	// need logout after change password
	logoutService := NewLogoutService(ctx, h.RequestContext)
	_, err = logoutService.Run(ctx, &user.Empty{})

	return &user.Empty{}, err
}
