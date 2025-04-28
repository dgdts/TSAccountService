package utils

import (
	"TSAccountService/biz/bizcontext"
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	tsredis "github.com/dgdts/ts-gobase/redis"
	redis "github.com/redis/go-redis/v9"
)

const (
	logoutPrefix = "TSAccountService:logout:token"

	defaultValue = ""
)

func GenerateLogoutKey(id string) string {
	return fmt.Sprintf("%s:%s", logoutPrefix, id)
}

func SetLogout(ctx *bizcontext.BizContext, token string, expiraMinute int) error {
	logoutKey := GenerateLogoutKey(token)
	expiration := time.Duration(expiraMinute) * time.Minute
	err := tsredis.GetConnection().SetEx(
		ctx,
		logoutKey,
		defaultValue,
		expiration,
	).Err()
	if err != nil {
		hlog.CtxErrorf(ctx, "SetLogout error:%v", err)
	}
	return err
}

func IsLogout(ctx context.Context, token string) bool {
	logoutKey := GenerateLogoutKey(token)
	err := tsredis.GetConnection().Get(ctx, logoutKey).Err()

	if err != nil {
		if errors.Is(err, redis.Nil) {
			return false
		} else {
			hlog.CtxErrorf(ctx, "IsLogout error:%v", err)
			return false
		}
	}
	return true
}
