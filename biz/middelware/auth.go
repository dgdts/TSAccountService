package middelware

import (
	"TSAccountService/biz/bizcontext"
	"TSAccountService/biz/constant"
	"TSAccountService/biz/errno"
	"TSAccountService/biz/response"
	"TSAccountService/biz/utils"
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	tsjwt "github.com/dgdts/ts-gobase/jwt"
)

var _ Middelware = (*auth)(nil)

type auth struct{}

func (a *auth) Init() {}

func (a *auth) GetOrder() int {
	return 1
}

func (a *auth) Name() string {
	return "auth"
}

var excludePath = map[string]struct{}{
	"/ping":                             {},
	"/api/v1/user/login-or-register":    {},
	"/api/v1/user/third-party-callback": {},
	"/api/v1/verify/send-sms":           {},
	"/api/v1/verify/send-email":         {},
}

func (a *auth) Do(ctx context.Context, c *app.RequestContext) {
	// check exclude path
	path := string(c.URI().Path())
	if _, ok := excludePath[path]; ok {
		return
	}
	// do auth
	token := c.Request.Header.Get(constant.HeaderToken)

	// check logout
	if utils.IsLogout(ctx, token) {
		response.JSONErr(c, errno.TokenLogoutErr)
		c.Abort()
		return
	}

	userMap, err := tsjwt.ValidateToken(token)
	if err != nil {
		hlog.CtxErrorf(ctx, "validate token error: %v", err)
		response.JSONErr(c, err)
		c.Abort()
		return
	}

	userID := userMap["ID"].(string)
	if userID == "" {
		hlog.CtxErrorf(ctx, "user id is empty")
		response.JSONErr(c, errno.TokenEmptyErr)
		c.Abort()
		return
	}

	bizCtx, ok := c.Get(constant.BizContext)
	if !ok {
		hlog.CtxErrorf(ctx, "get biz context error")
		response.JSONErr(c, errno.InternalErr)
		c.Abort()
		return
	}

	bizCtx.(*bizcontext.BizContext).User = &bizcontext.User{
		ID: userID,
	}

	c.Next(ctx)
}
