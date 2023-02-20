package config

import (
	"strings"

	"github.com/spf13/viper"
)

type Postgres struct {
	Host           string `mapstructure:"host"`
	Port           string `mapstructure:"port"`
	User           string `mapstructure:"user"`
	Password       string `mapstructure:"password"`
	DB             string `mapstructure:"db"`
	MigrationsPath string `mapstructure:"migrations-path"`
}

type Redis struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
}

type Mongo struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	DB       string `mapstructure:"db"`
}

type HTTPServer struct {
	IP   string `mapstructure:"ip"`
	Port string `mapstructure:"port"`
}

type Auth struct {
	SecretKey            string `mapstructure:"secret-key"`
	TokenExpirationHours uint   `mapstructure:"token-expiration-hours"`
}

type Config struct {
	Postgres   `mapstructure:"postgres"`
	Redis      `mapstructure:"redis"`
	Mongo      `mapstructure:"mongo"`
	HTTPServer `mapstructure:"http-server"`
	Auth       `mapstructure:"auth"`
}

func Read() (Config, error) {
	viper.AutomaticEnv()
	viper.SetEnvPrefix("MEDAD")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config/")

	if err := viper.ReadInConfig(); err != nil {
		return Config{}, err
	}

	var c Config
	if err := viper.Unmarshal(&c); err != nil {
		return Config{}, err
	}

	return c, nil
}
