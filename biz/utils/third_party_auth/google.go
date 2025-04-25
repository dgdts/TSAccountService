package third_party_auth

import (
	"TSAccountService/biz/bizcontext"
	"TSAccountService/biz/errno"
	"net/http"
	"sync"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var _ ThirdPartyOauth = (*googleAuth)(nil)

const (
	GoogleAuthType = "google"
)

var (
	gGoogleOauthConfig *oauth2.Config
	once               sync.Once
)

type GoogleAuthConfig struct {
	ClientID     string   `yaml:"client_id"`
	ClientSecret string   `yaml:"client_secret"`
	RedirectURL  string   `yaml:"redirect_url"`
	Scopes       []string `yaml:"scopes"`
}

func InitGoogleAuth(config GoogleAuthConfig) {
	once.Do(func() {
		gGoogleOauthConfig = &oauth2.Config{
			ClientID:     config.ClientID,
			ClientSecret: config.ClientSecret,
			RedirectURL:  config.RedirectURL,
			Scopes:       config.Scopes,
			Endpoint:     google.Endpoint,
		}
	})
}

type googleAuth struct {
}

func (g *googleAuth) Oauth(ctx *bizcontext.BizContext) error {
	if gGoogleOauthConfig == nil {
		hlog.Error("google auth config is nil")
		return errno.InternalErr
	}

	callbackUrl := gGoogleOauthConfig.AuthCodeURL("state")
	ctx.RequestContext.Redirect(http.StatusTemporaryRedirect, []byte(callbackUrl))
	ctx.RequestContext.Abort()
	return nil
}
