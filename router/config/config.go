package config

import (
	"github.com/lc-tut/club-portal/router/utils"
	"golang.org/x/oauth2"
)

type IConfig interface {
	ToAuthConfig() *AuthConfig
	ToMiddlewareConfig() *MiddlewareConfig
	ToV1Config() *V1Config
}

type Config struct {
	WhitelistUsers utils.WhitelistInfo
}

func (c *Config) ToAuthConfig() *AuthConfig {
	sessionCookie := NewSessionCookieOption()
	csrfCookie := NewCSRFCookieOption()
	googleConf := NewOAuth2Config()

	authConf := &AuthConfig{
		SessionCookieOptions: sessionCookie,
		CSRFCookieOptions:    csrfCookie,
		GoogleOAuthConfig:    googleConf,
		WhitelistUsers:       c.WhitelistUsers,
	}
	return authConf
}

func (c *Config) ToMiddlewareConfig() *MiddlewareConfig {
	middlewareConf := &MiddlewareConfig{
		WhitelistUsers: c.WhitelistUsers,
	}
	return middlewareConf
}

func (c *Config) ToV1Config() *V1Config {
	v1Conf := &V1Config{
		WhitelistUsers: c.WhitelistUsers,
	}
	return v1Conf
}

type AuthConfig struct {
	SessionCookieOptions *SessionCookieOption
	CSRFCookieOptions    *CSRFCookieOption
	GoogleOAuthConfig    *oauth2.Config
	WhitelistUsers       utils.WhitelistInfo
}

type MiddlewareConfig struct {
	WhitelistUsers utils.WhitelistInfo
}

type V1Config struct {
	WhitelistUsers utils.WhitelistInfo
}
