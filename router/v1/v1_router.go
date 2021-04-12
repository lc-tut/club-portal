package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/lc-tut/club-portal/repos"
	"github.com/lc-tut/club-portal/router/config"
	"github.com/lc-tut/club-portal/router/middlewares"
	"go.uber.org/zap"
)

type Handler struct {
	config *config.V1Config
	logger *zap.Logger
	repo   repos.IRepository
}

func newHandler(config *config.V1Config, logger *zap.Logger, repo repos.IRepository) *Handler {
	return &Handler{
		config: config,
		logger: logger,
		repo:   repo,
	}
}

type Router struct {
	rg         *gin.RouterGroup
	config     *config.V1Config
	logger     *zap.Logger
	repo       repos.IRepository
	middleware *middlewares.Middleware
}

func (r *Router) AddRouter() {
	h := newHandler(r.config, r.logger, r.repo)

	v1Group := r.rg.Group("/v1")
	{
		userGroup := v1Group.Group("/user")
		{
			userGroup.GET("/", h.GetUser())
			userGroup.GET("/:uuid", r.middleware.AdminOnly(), h.GetUserUUID())
		}
		clubGroup := v1Group.Group("/clubs")
		{
			clubGroup.GET("/", h.GetAllClub())
			clubGroup.POST("/", h.CreateClub())
			clubGroup.GET("/:clubslug", r.middleware.SetClubIDKey(), h.GetClub())
			clubGroup.PUT("/:clubslug", r.middleware.SetClubIDKey())
		}
	}
}

func NewV1Router(rg *gin.RouterGroup, config *config.V1Config, logger *zap.Logger, repo repos.IRepository, middleware *middlewares.Middleware) *Router {
	r := &Router{
		rg:         rg,
		config:     config,
		logger:     logger,
		repo:       repo,
		middleware: middleware,
	}
	return r
}
