package errno

var (
	OldPasswordErr    = NewErrNo(UserErrorCode+1, "old password error")
	UserNotFoundErr   = NewErrNo(UserErrorCode+2, "user not found")
	UpdateUserInfoErr = NewErrNo(UserErrorCode+3, "update user info error")
)
