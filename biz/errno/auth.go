package errno

var (
	ErrUnsupportedAuthType = NewErrNo(AuthErrorCode+1, "unsupported auth type")
)
