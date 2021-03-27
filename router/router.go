package router

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/lc-tut/club-portal/consts"
	"github.com/lc-tut/club-portal/router/auth"
	"github.com/lc-tut/club-portal/router/data"
	"go.uber.org/zap"
)

type IRouter interface {
	AddRouter()
}

func Init(engine *gin.Engine, logger *zap.Logger) error {
	config := data.NewConfig(true)

	store, err := data.NewRedisServer()

	if err != nil {
		return err
	}

	store.Options(*config.SessionCookieOptions)
	engine.Use(sessions.Sessions(consts.SessionCookieName, store))

	registerRouters(engine, config, logger)

	return nil
}

func registerRouters(engine *gin.Engine, config *data.Config, logger *zap.Logger) {
	apiGroup := engine.Group("/api")
	authRouter := auth.NewAuthRouter(apiGroup, config, logger)
	addRouter(authRouter)
}

func addRouter(r IRouter) {
	r.AddRouter()
}
