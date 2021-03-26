package router

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/lc-tut/club-portal/consts"
	"github.com/lc-tut/club-portal/router/auth"
	"github.com/lc-tut/club-portal/router/data"
)

type IRouter interface {
	AddRouter()
}

func Init(engine *gin.Engine) {
	config, err := data.NewConfig(true)

	if err != nil {
		panic(err)
	}

	store, err := data.NewRedisServer()

	if err != nil {
		panic(err)
	}

	store.Options(*config.SessionCookieOptions)
	engine.Use(sessions.Sessions(consts.SessionCookieName, store))

	registerRouters(engine, config)
}

func registerRouters(engine *gin.Engine, config *data.Config) {
	apiGroup := engine.Group("/api")
	authRouter := auth.NewAuthRouter(apiGroup, config)
	addRouter(authRouter)
}

func addRouter(r IRouter) {
	r.AddRouter()
}
