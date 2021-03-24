package router

import (
	"github.com/gin-gonic/gin"
	"github.com/lc-tut/club-portal/router/auth"
)

func Init(engine *gin.Engine) {
	apiGroup := engine.Group("/api")
	auth.AddAuthRouter(apiGroup)
}
