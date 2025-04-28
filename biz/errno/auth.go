package errno

var (
	UnsupportedAuthTypeErr      = NewErrNo(AuthErrorCode+1, "unsupported auth type")
	AccountOrPasswordErr        = NewErrNo(AuthErrorCode+2, "account or password error")
	VerifyCodeErr               = NewErrNo(AuthErrorCode+3, "verify code error")
	RequestTooManyVerifyCodeErr = NewErrNo(AuthErrorCode+4, "request too many verify code")

	TokenEmptyErr  = NewErrNo(AuthErrorCode+5, "token error")
	TokenLogoutErr = NewErrNo(AuthErrorCode+6, "token logout")
)
