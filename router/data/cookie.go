package data

import (
	"github.com/lc-tut/club-portal/consts"
	"github.com/lc-tut/club-portal/utils"
	"github.com/spf13/viper"
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

type SessionCookieOption struct {
	Path     string
	Domain   string
	MaxAge   int
	Secure   bool
	HttpOnly bool
	SameSite http.SameSite
}

func NewSessionCookieOption() *SessionCookieOption {
	domain := viper.GetString("domain")

	opt := &SessionCookieOption{
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

	return opt
}
