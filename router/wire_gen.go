// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package router

import (
	"github.com/lc-tut/club-portal/repos"
	"github.com/lc-tut/club-portal/router/config"
	"github.com/lc-tut/club-portal/router/utils"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// Injectors from router_wire.go:

func newServer(logger *zap.Logger, db *gorm.DB) (*Server, error) {
	options := config.NewSessionCookieOption()
	store, err := newRedisStore(options)
	if err != nil {
		return nil, err
	}
	engine := newGinEngine(logger, store)
	csrfCookieOption := config.NewCSRFCookieOption()
	oauth2Config := config.NewOAuth2Config()
	repository := repos.NewRepository(logger, db)
	whitelistInfo, err := utils.NewWhitelist(repository)
	if err != nil {
		return nil, err
	}
	configConfig := &config.Config{
		CSRFCookieOptions: csrfCookieOption,
		GoogleOAuthConfig: oauth2Config,
		WhitelistUsers:    whitelistInfo,
	}
	server := registerRouters(engine, configConfig, logger, repository)
	return server, nil
}