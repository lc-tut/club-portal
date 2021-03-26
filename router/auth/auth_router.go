package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/lc-tut/club-portal/router/data"
)

type Handler struct {
	config *data.Config
}

func newHandler(config *data.Config) *Handler {
	return &Handler{
		config: config,
	}
}

type Router struct {
	rg     *gin.RouterGroup
	config *data.Config
}

func (r *Router) AddRouter() {
	h := newHandler(r.config)

	authGroup := r.rg.Group("/auth")
	{
		authGroup.GET("/", h.Auth())
		authGroup.GET("/signin", h.SignIn())
		authGroup.GET("/callback", h.Callback())
		authGroup.POST("/destroy", h.Destroy())
	}
}

func NewAuthRouter(rg *gin.RouterGroup, config *data.Config) *Router {
	r := &Router{
		rg:     rg,
		config: config,
	}
	return r
}
