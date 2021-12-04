package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/lc-tut/club-portal/repos"
	"github.com/lc-tut/club-portal/router/config"
	"go.uber.org/zap"
)

type Handler struct {
	config *config.AuthConfig
	logger *zap.Logger
	repo   repos.IRepository
}

func newHandler(config *config.AuthConfig, logger *zap.Logger, repo repos.IRepository) *Handler {
	return &Handler{
		config: config,
		repo:   repo,
		logger: logger,
	}
}

type Router struct {
	rg     *gin.RouterGroup
	config *config.AuthConfig
	logger *zap.Logger
	repo   repos.IRepository
}

func (r *Router) AddRouter() {
	r.logger.Debug("initializing auth router")
	h := newHandler(r.config, r.logger, r.repo)

	authGroup := r.rg.Group("/auth")
	{
		authGroup.GET("/", h.Auth())
		authGroup.GET("/signin", h.SignIn())
		authGroup.GET("/callback", h.Callback())
		authGroup.POST("/destroy", h.Destroy())
	}
}

func NewAuthRouter(rg *gin.RouterGroup, config *config.AuthConfig, logger *zap.Logger, repo repos.IRepository) *Router {
	r := &Router{
		rg:     rg,
		config: config,
		logger: logger,
		repo:   repo,
	}
	return r
}
