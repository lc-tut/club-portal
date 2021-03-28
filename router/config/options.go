package config

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/lc-tut/club-portal/consts"
	"github.com/lc-tut/club-portal/utils"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"net/http"
)

type CSRFCookieOption struct {
	Path     string
	Domain   string
	MaxAge   int
	Secure   bool
	HttpOnly bool
	SameSite http.SameSite
}

func NewCSRFCookieOption() *CSRFCookieOption {
	domain := viper.GetString("domain")

	opt := &CSRFCookieOption{
		Path:     consts.CookiePath,
		Domain:   domain,
		MaxAge:   consts.CookieCSRFMaxAge,
		HttpOnly: consts.CookieHttpOnly,
	}

	if utils.IsLocal() {
		opt.Secure = false
		opt.SameSite = http.SameSiteLaxMode
	} else {
		opt.Secure = true
		opt.SameSite = http.SameSiteStrictMode
	}

	return opt
}

func NewSessionCookieOption() sessions.Options {
	domain := viper.GetString("domain")

	opt := &sessions.Options{
		Path:     consts.CookiePath,
		Domain:   domain,
		MaxAge:   consts.CookieSessionMaxAge,
		HttpOnly: consts.CookieHttpOnly,
		SameSite: http.SameSiteLaxMode,
	}

	if utils.IsLocal() {
		opt.Secure = false
	} else {
		opt.Secure = true
	}

	return *opt
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
