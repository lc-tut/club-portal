package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/lc-tut/club-portal/router/config"
	"github.com/lc-tut/club-portal/router/middlewares"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Handler struct {
	config *config.V1Config
	logger *zap.Logger
	db     *gorm.DB
}

func newHandler(config *config.V1Config, logger *zap.Logger, db *gorm.DB) *Handler {
	return &Handler{
		config: config,
		logger: logger,
		db:     db,
	}
}

type Router struct {
	rg         *gin.RouterGroup
	config     *config.V1Config
	logger     *zap.Logger
	db         *gorm.DB
	middleware *middlewares.Middleware
}

func (r *Router) AddRouter() {
	h := newHandler(r.config, r.logger, r.db)

	v1Group := r.rg.Group("/v1", r.middleware.CheckSession())
	{
		userGroup := v1Group.Group("/user")
		{
			userGroup.GET("/", h.GetUser())
			userGroup.GET("/:uuid", h.GetUserUUID(), r.middleware.AdminOnly())
		}
	}
}

func NewV1Router(rg *gin.RouterGroup, config *config.V1Config, logger *zap.Logger, db *gorm.DB, middleware *middlewares.Middleware) *Router {
	r := &Router{
		rg:         rg,
		config:     config,
		logger:     logger,
		db:         db,
		middleware: middleware,
	}
	return r
}
