package middlewares

import (
	"github.com/lc-tut/club-portal/repos/users"
	"github.com/lc-tut/club-portal/router/config"
	"go.uber.org/zap"
)

type Middleware struct {
	config *config.MiddlewareConfig
	logger *zap.Logger
	repo   users.UserRepo
}

func NewMiddleware(config *config.MiddlewareConfig, logger *zap.Logger, repo users.UserRepo) *Middleware {
	mw := &Middleware{
		config: config,
		logger: logger,
		repo:   repo,
	}
	return mw
}
