// +build wireinject

package router

import (
	"github.com/google/wire"
	"github.com/lc-tut/club-portal/router/data"
	"go.uber.org/zap"
)

func newServer(logger *zap.Logger) (*Server, error) {
	wire.Build(
		newGinEngine,
		registerRouters,
		newRedisServer,
		data.NewWhitelist,
		data.NewCSRFCookieOption,
		data.NewSessionCookieOption,
		data.NewOAuth2Config,
		wire.Struct(new(data.Config), "*"),
	)

	return nil, nil
}
