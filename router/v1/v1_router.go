package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/lc-tut/club-portal/router/data"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Handler struct {
	config *data.Config
	logger *zap.Logger
	db     *gorm.DB
}

func newHandler(config *data.Config, logger *zap.Logger, db *gorm.DB) *Handler {
	return &Handler{
		config: config,
		logger: logger,
		db:     db,
	}
}

type Router struct {
	rg     *gin.RouterGroup
	config *data.Config
	logger *zap.Logger
	db     *gorm.DB
}

func (r *Router) AddRouter() {

}

func NewV1Router(rg *gin.RouterGroup, config *data.Config, logger *zap.Logger, db *gorm.DB) *Router {
	r := &Router{
		rg:     rg,
		config: config,
		logger: logger,
		db:     db,
	}
	return r
}
