package verify_code

import (
	"TSAccountService/biz/errno"
	"context"
	"errors"
	"sync"

	"github.com/dgdts/ts-gobase/verify_code"
	tsverify "github.com/dgdts/ts-gobase/verify_code"
	"github.com/redis/go-redis/v9"
)

var gRedisVerifyCodeService tsverify.VerifyCode
var once sync.Once

func InitRedisVerifyCodeService(rdb redis.UniversalClient, opt tsverify.VerifyCodeOption, template string) {
	once.Do(func() {
		gRedisVerifyCodeService = tsverify.NewRedisVerifyCodeService(
			rdb,
			opt,
			template,
		)
	})
}

func StoreVerifyCode(ctx context.Context, key, code, ip string) error {
	if gRedisVerifyCodeService == nil {
		return errors.New("redis verify code service not init")
	}
	_, err := gRedisVerifyCodeService.StoreVerifyCode(ctx, key, code, ip)

	if errors.Is(err, verify_code.RedisStoreVerifyCodeErrLimit) {
		return errno.RequestTooManyVerifyCodeErr
	}

	return err
}

func ValidateVerifyCode(ctx context.Context, key, code, ip string) error {
	if gRedisVerifyCodeService == nil {
		return errors.New("redis verify code service not init")
	}
	return gRedisVerifyCodeService.ValidateVerifyCode(ctx, key, code, ip)
}
