package middlewares

import (
	"github.com/lc-tut/club-portal/repos"
	"github.com/lc-tut/club-portal/router/config"
	"go.uber.org/zap"
)

type Middleware struct {
	config *config.MiddlewareConfig
	logger *zap.Logger
	repo   repos.IRepository
}

func NewMiddleware(config *config.MiddlewareConfig, logger *zap.Logger, repo repos.IRepository) *Middleware {
	mw := &Middleware{
		config: config,
		logger: logger,
		repo:   repo,
	}
	return mw
}
