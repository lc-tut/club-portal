// +build wireinject

package router

import (
	"github.com/google/wire"
	"github.com/lc-tut/club-portal/repos"
	"github.com/lc-tut/club-portal/repos/users"
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
		repos.NewRepository,
		config.NewCSRFCookieOption,
		config.NewOAuth2Config,
		config.NewSessionCookieOption,
		utils.NewWhitelist,
		wire.Struct(new(config.Config), "*"),
		wire.Bind(new(config.IConfig), new(*config.Config)),
		wire.Bind(new(repos.IRepository), new(*repos.Repository)),
		wire.Bind(new(users.UserRepo), new(*repos.Repository)),
	)

	return nil, nil
}
