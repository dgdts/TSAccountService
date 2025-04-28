package verify

import (
	"context"
	"testing"

	"TSAccountService/biz/bizcontext"
	verify "TSAccountService/hertz_gen/verify"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
)

func TestSendEmallService_Run(t *testing.T) {
	ctx := context.Background()
	c := app.NewContext(1)
	s := NewSendEmallService(ctx, c)
	// init req and assert value
	req := &verify.SendEmallReq{}
	bizctx := &bizcontext.BizContext{}
	resp, err := s.Run(bizctx, req)
	assert.DeepEqual(t, nil, resp)
	assert.DeepEqual(t, nil, err)
	// todo edit your unit test.
}
