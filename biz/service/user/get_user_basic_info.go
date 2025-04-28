package user

import (
	"context"

	"TSAccountService/biz/bizcontext"
	"TSAccountService/biz/errno"
	"TSAccountService/biz/model"
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
	userBasic, err := model.FindUserByID(ctx.User.ID)
	if err != nil {
		return nil, errno.UserNotFoundErr
	}

	resp = &user.GetUserBasicInfoResp{
		Nickname:   userBasic.Nickname,
		UserID:     userBasic.ID.Hex(),
		Avatar:     userBasic.Avatar,
		CreateTime: userBasic.CreatedAt.Unix(),
		UpdateTime: userBasic.UpdatedAt.Unix(),
	}

	return resp, nil
}
