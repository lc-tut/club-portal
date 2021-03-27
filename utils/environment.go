package utils

import "github.com/spf13/viper"

func IsLocal() bool {
	mode := viper.GetString("mode")
	return mode == "local"
}

func IsDev() bool {
	mode := viper.GetString("mode")
	return mode == "development"
}

func IsProd() bool {
	mode := viper.GetString("mode")
	return mode == "production"
}
