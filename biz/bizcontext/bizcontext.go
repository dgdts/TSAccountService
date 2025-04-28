package bizcontext

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
)

type User struct {
	ID string
}

type BizContext struct {
	context.Context

	RequestContext *app.RequestContext
	User           *User
}
