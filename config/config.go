package config

import (
	"github.com/spf13/viper"
)

type Postgres struct {
	Host     string
	Port     string
	User     string
	Password string
}

type Redis struct {
	Host     string
	Port     string
	Password string
}

type Config struct {
	Postgres
	Redis
}

func Read() (*Config, error) {
	viper.AutomaticEnv()
	viper.SetEnvPrefix("MEDAD")

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config/")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var c Config
	if err := viper.Unmarshal(&c); err != nil {
		return nil, err
	}

	return &c, nil
}
