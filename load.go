package main

import (
	"github.com/lc-tut/club-portal/router"
	"github.com/lc-tut/club-portal/utils"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func loadConfig() error {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath("/run/secrets") // for docker secrets

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	viper.SetDefault("mode", "local")
	viper.SetDefault("domain", "localhost")

	return nil
}

func newZapLogger() (*zap.Logger, error) {
	if utils.IsProd() {
		return zap.NewProduction()
	} else {
		return zap.NewDevelopment()
	}
}

func newServer() (*router.Server, error) {
	if err := loadConfig(); err != nil {
		return nil, err
	}

	logger, err := newZapLogger()

	if err != nil {
		return nil, err
	}

	server, err := router.NewServer(logger)

	if err != nil {
		return nil, err
	}

	return server, nil
}
