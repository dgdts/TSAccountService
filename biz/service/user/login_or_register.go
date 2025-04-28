package user

import (
	"context"
	"errors"

	"TSAccountService/biz/bizcontext"
	"TSAccountService/biz/errno"
	"TSAccountService/biz/model"
	"TSAccountService/biz/utils/third_party_auth"
	user "TSAccountService/hertz_gen/user"
	"TSAccountService/kit/verify_code"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"go.mongodb.org/mongo-driver/mongo"

	tsjwt "github.com/dgdts/ts-gobase/jwt"
)

type LoginOrRegisterService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewLoginOrRegisterService(Context context.Context, RequestContext *app.RequestContext) *LoginOrRegisterService {
	return &LoginOrRegisterService{RequestContext: RequestContext, Context: Context}
}

func verifySMSCode(phone, code, ip string) bool {
	return verify_code.ValidateVerifyCode(context.Background(), phone, code, ip) == nil
}

func verifyEmallCode(email, code, ip string) bool {
	return verify_code.ValidateVerifyCode(context.Background(), email, code, ip) == nil
}

func (h *LoginOrRegisterService) Run(ctx *bizcontext.BizContext, req *user.LoginOrRegisterReq) (resp *user.LoginOrRegisterResp, err error) {
	if req.AuthType != "" {
		return nil, third_party_auth.Oauth(ctx, req.AuthType)
	}

	if req.Phone != "" {
		return h.loginOrRegisterWithPhone(ctx, req)
	} else if req.Email != "" {
		return h.loginOrRegisterWithEmail(ctx, req)
	}

	return nil, errno.ParameterErr
}

func (h *LoginOrRegisterService) loginOrRegisterWithPhone(ctx *bizcontext.BizContext, req *user.LoginOrRegisterReq) (resp *user.LoginOrRegisterResp, err error) {
	var userBasic *model.UserBasic
	if req.VerifyCode != "" {
		if verifySMSCode(req.Phone, req.VerifyCode, h.RequestContext.ClientIP()) {
			userBasic, err = model.FindUserByPhone(req.Phone)
			if errors.Is(err, mongo.ErrNoDocuments) {
				userBasic = &model.UserBasic{
					Phone:    req.Phone,
					Nickname: req.Phone,
				}
				err = model.CreateUser(userBasic)
				if err != nil {
					hlog.CtxErrorf(ctx, "create user failed, err: %v", err)
					return nil, errno.InternalErr
				}
			} else if err != nil {
				hlog.CtxErrorf(ctx, "find user failed, err: %v", err)
				return nil, errno.InternalErr
			}
		} else {
			return nil, errno.VerifyCodeErr
		}
	} else if req.Password != "" {
		userBasic, err = model.FindUserByPhoneAndPassword(req.Phone, req.Password)
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errno.AccountOrPasswordErr
		} else if err != nil {
			hlog.CtxErrorf(ctx, "find user failed, err: %v", err)
			return nil, errno.InternalErr
		}
	}

	// auth success
	claims := map[string]any{
		"ID": userBasic.ID.Hex(),
	}

	token, err := tsjwt.GenerateToken(claims)
	if err != nil {
		hlog.CtxErrorf(ctx, "generate token failed, err: %v", err)
		return nil, errno.InternalErr
	}

	resp = &user.LoginOrRegisterResp{
		Token: token,
	}

	return resp, nil
}

func (h *LoginOrRegisterService) loginOrRegisterWithEmail(ctx *bizcontext.BizContext, req *user.LoginOrRegisterReq) (resp *user.LoginOrRegisterResp, err error) {
	var userBasic *model.UserBasic
	if req.VerifyCode != "" {
		if verifyEmallCode(req.Email, req.VerifyCode, h.RequestContext.ClientIP()) {
			userBasic, err = model.FindUserByEmail(req.Email)
			if errors.Is(err, mongo.ErrNoDocuments) {
				userBasic = &model.UserBasic{
					Email:    req.Email,
					Nickname: req.Email,
				}
				err = model.CreateUser(userBasic)
				if err != nil {
					hlog.CtxErrorf(ctx, "create user failed, err: %v", err)
					return nil, errno.InternalErr
				}
			} else if err != nil {
				hlog.CtxErrorf(ctx, "find user failed, err: %v", err)
				return nil, errno.InternalErr
			}
		} else {
			return nil, errno.VerifyCodeErr
		}
	} else if req.Password != "" {
		userBasic, err = model.FindUserByEmailAndPassword(req.Email, req.Password)
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errno.AccountOrPasswordErr
		} else if err != nil {
			hlog.CtxErrorf(ctx, "find user failed, err: %v", err)
			return nil, errno.InternalErr
		}
	}

	// auth success
	claims := map[string]any{
		"ID": userBasic.ID.Hex(),
	}
	token, err := tsjwt.GenerateToken(claims)
	if err != nil {
		hlog.CtxErrorf(ctx, "generate token failed, err: %v", err)
		return nil, errno.InternalErr
	}
	resp = &user.LoginOrRegisterResp{
		Token: token,
	}

	return resp, nil
}
