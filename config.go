package main

import (
	"github.com/lc-tut/club-portal/consts"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

func newZapProdConfig() (*zap.Logger, error) {
	prodEncoderConfig := zap.NewProductionEncoderConfig()
	prodEncoderConfig.EncodeTime = jstTimeEncoder
	prodEncoderConfig.EncodeDuration = zapcore.StringDurationEncoder

	conf := zap.NewProductionConfig()
	conf.DisableCaller = true
	conf.EncoderConfig = prodEncoderConfig

	return conf.Build()
}

func newZapDevConfig() (*zap.Logger, error) {
	devEncoderConfig := zap.NewDevelopmentEncoderConfig()
	devEncoderConfig.EncodeTime = jstTimeEncoder

	conf := zap.NewDevelopmentConfig()
	conf.EncoderConfig = devEncoderConfig

	return conf.Build()
}

func jstTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.In(consts.JST).Format(time.RFC3339))
}
