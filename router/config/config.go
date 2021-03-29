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
	CSRFCookieOptions *CSRFCookieOption
	GoogleOAuthConfig *oauth2.Config
	WhitelistUsers    utils.WhitelistInfo
}

func (c *Config) ToAuthConfig() *AuthConfig {
	authConf := &AuthConfig{
		CSRFCookieOptions: c.CSRFCookieOptions,
		GoogleOAuthConfig: c.GoogleOAuthConfig,
		WhitelistUsers:    c.WhitelistUsers,
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
	CSRFCookieOptions *CSRFCookieOption
	GoogleOAuthConfig *oauth2.Config
	WhitelistUsers    utils.WhitelistInfo
}

type MiddlewareConfig struct {
	WhitelistUsers utils.WhitelistInfo
}

type V1Config struct {
	WhitelistUsers utils.WhitelistInfo
}
