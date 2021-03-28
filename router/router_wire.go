// +build wireinject

package router

import (
	"github.com/google/wire"
	"github.com/lc-tut/club-portal/router/config"
	"github.com/lc-tut/club-portal/router/utils"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func newServer(logger *zap.Logger, db *gorm.DB) (*Server, error) {
	wire.Build(
		newGinEngine,
		registerRouters,
		newRedisStore,
		utils.NewWhitelist,
		wire.Struct(new(config.Config), "*"),
		wire.Bind(new(config.IConfig), new(*config.Config)),
	)

	return nil, nil
}