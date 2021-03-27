package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/lc-tut/club-portal/router/data"
	"go.uber.org/zap"
)

type Handler struct {
	config *data.Config
	logger *zap.Logger
}

func newHandler(config *data.Config, logger *zap.Logger) *Handler {
	return &Handler{
		config: config,
		logger: logger,
	}
}

type Router struct {
	rg     *gin.RouterGroup
	config *data.Config
	logger *zap.Logger
}

func (r *Router) AddRouter() {
	h := newHandler(r.config, r.logger)

	authGroup := r.rg.Group("/auth")
	{
		authGroup.GET("/", h.Auth())
		authGroup.GET("/signin", h.SignIn())
		authGroup.GET("/callback", h.Callback())
		authGroup.POST("/destroy", h.Destroy())
	}
}

func NewAuthRouter(rg *gin.RouterGroup, config *data.Config, logger *zap.Logger) *Router {
	r := &Router{
		rg:     rg,
		config: config,
		logger: logger,
	}
	return r
}
