package config

import (
	"strings"

	"github.com/spf13/viper"
)

type Postgres struct {
	Host     string
	Port     string
	User     string
	Password string
	DB       string
}

type Redis struct {
	Host     string
	Port     string
	User     string
	Password string
}

type Mongo struct {
	Host     string
	Port     string
	User     string
	Password string
	DB       string
}

type HTTPServer struct {
	IP   string
	Port string
}

type Config struct {
	Postgres
	Redis
	Mongo
	HTTPServer
}

func Read() (*Config, error) {
	viper.AutomaticEnv()
	viper.SetEnvPrefix("MEDAD")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))

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
