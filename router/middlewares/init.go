package middlewares

import (
	"github.com/lc-tut/club-portal/router/config"
	"go.uber.org/zap"
)

type Middleware struct {
	config *config.MiddlewareConfig
	logger *zap.Logger
}

func NewMiddleware(config *config.MiddlewareConfig, logger *zap.Logger) *Middleware {
	mw := &Middleware{
		config: config,
		logger: logger,
	}
	return mw
}
