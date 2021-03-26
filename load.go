package main

import "github.com/spf13/viper"

func loadConfig() error {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath("/run/secrets") // for docker secrets

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	return nil
}
