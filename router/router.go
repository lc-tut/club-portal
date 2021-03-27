package router

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/lc-tut/club-portal/consts"
	"github.com/lc-tut/club-portal/router/auth"
	"github.com/lc-tut/club-portal/router/data"
	"github.com/lc-tut/club-portal/utils"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"time"
)

type IRouter interface {
	AddRouter()
}

type Server struct {
	*gin.Engine
}

func NewServer(logger *zap.Logger) (*Server, error) {
	server, err := newServer(logger)

	if err != nil {
		return nil, err
	}

	return server, nil
}

func newGinEngine(logger *zap.Logger, ss redis.Store) *gin.Engine {
	engine := gin.New()

	engine.Use(ginzap.Ginzap(logger, time.RFC3339, false))
	engine.Use(ginzap.RecoveryWithZap(logger, !utils.IsProd()))
	engine.Use(sessions.Sessions(consts.SessionCookieName, ss))

	return engine
}

func registerRouters(engine *gin.Engine, config *data.Config, logger *zap.Logger) *Server {
	apiGroup := engine.Group("/api")
	authRouter := auth.NewAuthRouter(apiGroup, config, logger)
	addRouter(authRouter)

	return &Server{engine}
}

func addRouter(r IRouter) {
	r.AddRouter()
}

func newRedisServer() (redis.Store, error) {
	secretKey := viper.GetString("redis_secret")
	store, err := redis.NewStore(10, "tcp", "redis:6379", viper.GetString("redis_password"), []byte(secretKey))

	if err != nil {
		return nil, err
	}

	return store, nil
}
