package router

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/lc-tut/club-portal/consts"
	"github.com/lc-tut/club-portal/repos"
	"github.com/lc-tut/club-portal/router/auth"
	"github.com/lc-tut/club-portal/router/config"
	"github.com/lc-tut/club-portal/router/middlewares"
	v1 "github.com/lc-tut/club-portal/router/v1"
	"github.com/lc-tut/club-portal/utils"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type IRouter interface {
	AddRouter()
}

type Server struct {
	*gin.Engine
}

func newRedisStore(opt sessions.Options) (redis.Store, error) {
	secretKey := viper.GetString("redis_secret")
	address := fmt.Sprintf("%s:%s", viper.GetString("redis_address"), viper.GetString("redis_port"))
	pass := viper.GetString("redis_password")
	store, err := redis.NewStore(10, "tcp", address, pass, []byte(secretKey))

	if err != nil {
		return nil, err
	}

	store.Options(opt)

	return store, nil
}

func newGinEngine(logger *zap.Logger, ss redis.Store) *gin.Engine {
	logger.Debug("initializing gin engine")
	if utils.IsProd() {
		gin.SetMode(gin.ReleaseMode)
	}

	engine := gin.New()

	engine.Use(middlewares.LoggerMiddleware(logger))
	engine.Use(ginzap.RecoveryWithZap(logger, !utils.IsProd()))
	engine.Use(sessions.Sessions(consts.SessionCookieName, ss))

	engine.MaxMultipartMemory = consts.UploadSize

	return engine
}

func registerRouters(engine *gin.Engine, config config.IConfig, logger *zap.Logger, repo repos.IRepository) *Server {
	logger.Debug("initializing registerRouter")

	logger.Debug("initializing middleware")
	mw := middlewares.NewMiddleware(config.ToMiddlewareConfig(), logger, repo)

	apiGroup := engine.Group("/api")

	authRouter := auth.NewAuthRouter(apiGroup, config.ToAuthConfig(), logger, repo)
	v1Router := v1.NewV1Router(apiGroup, config.ToV1Config(), logger, repo, mw)

	addRouter(authRouter, v1Router)

	return &Server{engine}
}

func addRouter(routers ...IRouter) {
	for _, r := range routers {
		r.AddRouter()
	}
}

func NewServer(logger *zap.Logger, db *gorm.DB) (*Server, error) {
	logger.Debug("initializing server")
	server, err := newServer(logger, db)

	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	return server, nil
}
