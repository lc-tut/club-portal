package admins

import (
	"github.com/lc-tut/club-portal/repos"
	"github.com/lc-tut/club-portal/router/config"
	"go.uber.org/zap"
)

type handler struct {
	config *config.V1Config
	logger *zap.Logger
	repo   repos.IRepository
}

func NewAdminHandler(config *config.V1Config, logger *zap.Logger, repo repos.IRepository) *handler {
	h := &handler{
		config: config,
		logger: logger,
		repo:   repo,
	}
	return h
}
