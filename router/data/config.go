package data

import (
	"fmt"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type Config struct {
	SessionCookieOptions *SessionCookieOption
	CSRFCookieOptions    *CSRFCookieOption
	GoogleOAuthConfig    *oauth2.Config
	WhitelistUsers       WhitelistInfo
}

func NewOAuth2Config() (conf *oauth2.Config) {
	domain := viper.GetString("domain")
	redirectURL := fmt.Sprintf("http://%s:8080/api/auth/callback", domain)

	conf = &oauth2.Config{
		ClientID:     viper.GetString("client_id"),
		ClientSecret: viper.GetString("client_secret"),
		Endpoint:     google.Endpoint,
		RedirectURL:  redirectURL,
		Scopes:       []string{"profile", "email"},
	}
	return
}
