package main

import (
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/lc-tut/club-portal/router"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"time"
)

func loadConfig() error {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath("/run/secrets") // for docker secrets

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	viper.SetDefault("production", false)

	return nil
}

func isDev() bool {
	mode := viper.GetBool("production")
	return !mode
}

func newZapLogger() (*zap.Logger, error) {
	if isDev() {
		return zap.NewDevelopment()
	} else {
		return zap.NewProduction()
	}
}

func newGinEngine(logger *zap.Logger) (*gin.Engine, error) {
	if err := loadConfig(); err != nil {
		return nil, err
	}

	engine := gin.New()

	engine.Use(ginzap.Ginzap(logger, time.RFC3339, false))
	engine.Use(ginzap.RecoveryWithZap(logger, isDev()))

	err := router.Init(engine, logger)

	if err != nil {
		return nil, err
	}

	return engine, nil
}
