package third_party_auth

import (
	"TSAccountService/biz/bizcontext"
	"TSAccountService/biz/errno"
)

type ThirdPartyOauth interface {
	Oauth(ctx *bizcontext.BizContext) error
}

func Oauth(ctx *bizcontext.BizContext, authType string) error {
	var auth ThirdPartyOauth
	switch authType {
	case GoogleAuthType:
		auth = new(googleAuth)
	default:
		return errno.ErrUnsupportedAuthType
	}
	return auth.Oauth(ctx)
}
