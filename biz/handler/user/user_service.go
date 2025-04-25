package user

import (
	"context"

	"TSAccountService/biz/response"
	logic "TSAccountService/biz/service/user"
	_ "TSAccountService/hertz_gen/user"

	"github.com/cloudwego/hertz/pkg/app"
)

// LoginOrRegister .
// @router /api/v1/user/login-or-register [POST]
func LoginOrRegister(ctx context.Context, c *app.RequestContext) {
	response.JSON(ctx, c, logic.NewLoginOrRegisterService(ctx, c).Run)
}

// GetUserBasicInfo .
// @router /api/v1/user/get-user-basic_info [POST]
func GetUserBasicInfo(ctx context.Context, c *app.RequestContext) {
	response.JSON(ctx, c, logic.NewGetUserBasicInfoService(ctx, c).Run)
}

// UpdateUserInfo .
// @router /api/v1/user/update-info [POST]
func UpdateUserInfo(ctx context.Context, c *app.RequestContext) {
	response.JSON(ctx, c, logic.NewUpdateUserInfoService(ctx, c).Run)
}

// ChangePassword .
// @router /api/v1/user/change-password [POST]
func ChangePassword(ctx context.Context, c *app.RequestContext) {
	response.JSON(ctx, c, logic.NewChangePasswordService(ctx, c).Run)
}

// Logout .
// @router /api/v1/user/logout [POST]
func Logout(ctx context.Context, c *app.RequestContext) {
	response.JSON(ctx, c, logic.NewLogoutService(ctx, c).Run)
}

// CheckPasswordSet .
// @router /api/v1/user/check-password-set [POST]
func CheckPasswordSet(ctx context.Context, c *app.RequestContext) {
	response.JSON(ctx, c, logic.NewCheckPasswordSetService(ctx, c).Run)
}

// RegisterWithThirdParty .
// @router /api/v1/user/register-with-third-party [POST]
func RegisterWithThirdParty(ctx context.Context, c *app.RequestContext) {
	response.JSON(ctx, c, logic.NewRegisterWithThirdPartyService(ctx, c).Run)
}

// ThirdPartyCallback .
// @router /api/v1/user/third-party/callback [GET]
func ThirdPartyCallback(ctx context.Context, c *app.RequestContext) {
	response.JSON(ctx, c, logic.NewThirdPartyCallbackService(ctx, c).Run)
}
