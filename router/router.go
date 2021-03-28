package router

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/lc-tut/club-portal/consts"
	"github.com/lc-tut/club-portal/router/auth"
	"github.com/lc-tut/club-portal/router/config"
	"github.com/lc-tut/club-portal/router/middlewares"
	v1 "github.com/lc-tut/club-portal/router/v1"
	"github.com/lc-tut/club-portal/utils"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"time"
)

type IRouter interface {
	AddRouter()
}

type Server struct {
	*gin.Engine
}

func NewServer(logger *zap.Logger, db *gorm.DB) (*Server, error) {
	server, err := newServer(logger, db)

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

func registerRouters(engine *gin.Engine, config config.IConfig, logger *zap.Logger, db *gorm.DB) *Server {
	mw := middlewares.NewMiddleware(config.ToMiddlewareConfig(), logger)

	apiGroup := engine.Group("/api")

	authRouter := auth.NewAuthRouter(apiGroup, config.ToAuthConfig(), logger)
	v1Router := v1.NewV1Router(apiGroup, config.ToV1Config(), logger, db, mw)

	addRouter(authRouter)
	addRouter(v1Router)

	return &Server{engine}
}

func addRouter(r IRouter) {
	r.AddRouter()
}

func newRedisStore() (redis.Store, error) {
	secretKey := viper.GetString("redis_secret")
	store, err := redis.NewStore(10, "tcp", "redis:6379", viper.GetString("redis_password"), []byte(secretKey))

	if err != nil {
		return nil, err
	}

	return store, nil
}
