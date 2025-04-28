package biz_config

import (
	"TSAccountService/biz/utils"
	"TSAccountService/biz/utils/third_party_auth"
	"os"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"gopkg.in/yaml.v3"

	"TSAccountService/biz/config/global_config"

	tsjwt "github.com/dgdts/ts-gobase/jwt"
	tsmongodb "github.com/dgdts/ts-gobase/mongo"
	tsredis "github.com/dgdts/ts-gobase/redis"
	tsverify "github.com/dgdts/ts-gobase/verify_code"
)

var bizConfig *BizConfig

type MockConfig struct {
	Enable    bool   `yaml:"enable"`
	SMSCode   string `yaml:"sms_code"`
	EmailCode string `yaml:"email_code"`
}

type BizConfig struct {
	GoogleAuthConfig third_party_auth.GoogleAuthConfig `yaml:"google_auth_config"`
	MongoDBConfig    tsmongodb.MongoClientConfig       `yaml:"mongo_db_config"`
	JWTConfig        tsjwt.JWTConfig                   `yaml:"jwt_config"`
	RedisConfig      map[string]*tsredis.RedisClient   `yaml:"redis_config"`
	VerifyCodeOption tsverify.VerifyCodeOption         `yaml:"verify_code_option"`

	MockConfig MockConfig `yaml:"mock_config"`
}

func BizConfigInit(config *global_config.Config) {
	switch config.Mode {
	case global_config.ModeTypeNacos:
		bizConfigNacosInit(config)
	default:
		bizConfigLocalInit(config)
	}
}

func bizConfigLocalInit(config *global_config.Config) {
	configBytes, err := os.ReadFile(config.LocalConfigPath)
	if err != nil {
		panic(err)
	}

	var tmpConfig BizConfig
	err = yaml.Unmarshal(configBytes, &tmpConfig)
	if err != nil {
		panic(err)
	}
	setBizConfig(&tmpConfig)
	hlog.Infof("local config: %v", GetBizConfig())
}

func bizConfigNacosInit(config *global_config.Config) {
	utils.InitRemoteConfig(*config.RemoteConfig)

	var tmpConfig BizConfig
	err := utils.GetRemoteConfig(config.ConfigDataID, config.ConfigGroup, &tmpConfig)
	if err != nil {
		panic(err)
	}

	setBizConfig(&tmpConfig)
	hlog.Infof("remote config: %v", GetBizConfig())
	// could add watch for hot update, but we should consider the race condition when update config
	// go watch(config.ConfigGroup, config.ConfigDataID)
}

func setBizConfig(config *BizConfig) {
	bizConfig = config
}

func GetBizConfig() *BizConfig {
	return bizConfig
}

func watch(group string, key string) {
	c, err := utils.WatchRemoteConfig(group, key)
	if err != nil {
		hlog.Errorf("watch failed. error: %s", err.Error())
		panic(err)
	}

	for resp := range c {
		tmpConfig := &BizConfig{}
		err := yaml.Unmarshal([]byte(resp), &tmpConfig)
		if err != nil {
			panic(err)
		}
		hlog.Infof("remote config Change tmp: %v", tmpConfig)
		setBizConfig(tmpConfig)
	}
	hlog.Infof("watch config end")
}
