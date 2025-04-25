package bizcontext

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
)

type User struct {
	UserName string `json:"user_name"`
}

type BizContext struct {
	context.Context

	RequestContext *app.RequestContext
	User           *User
}
