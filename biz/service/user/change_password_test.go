package user

import (
	"context"
	"testing"

	"TSAccountService/biz/bizcontext"
	user "TSAccountService/hertz_gen/user"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
)

func TestChangePasswordService_Run(t *testing.T) {
	ctx := context.Background()
	c := app.NewContext(1)
	s := NewChangePasswordService(ctx, c)
	// init req and assert value
	req := &user.ChangePasswordReq{}
	bizctx := &bizcontext.BizContext{}
	resp, err := s.Run(bizctx, req)
	assert.DeepEqual(t, nil, resp)
	assert.DeepEqual(t, nil, err)
	// todo edit your unit test.
}
