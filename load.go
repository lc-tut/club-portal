package main

import (
	"fmt"
	"github.com/lc-tut/club-portal/router"
	"github.com/lc-tut/club-portal/utils"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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

func newDB() (*gorm.DB, error) {
	dbuser := viper.GetString("mariadb_user")
	dbpass := viper.GetString("mariadb_password")
	dbname := viper.GetString("mariadb_dbname")
	dsn := fmt.Sprintf("%s:%s@tcp(mariadb:3306)/%s?charset=utf8mb4&parseTime=True&loc=Asia%%2FTokyo", dbuser, dbpass, dbname)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	if utils.IsLocal() {
		db = db.Debug()
	}

	return db, err
}

func newServer() (*router.Server, error) {
	if err := loadConfig(); err != nil {
		return nil, err
	}

	logger, err := newZapLogger()

	if err != nil {
		return nil, err
	}

	db, err := newDB()

	if err != nil {
		return nil, err
	}

	server, err := router.NewServer(logger, db)

	if err != nil {
		return nil, err
	}

	return server, nil
}
