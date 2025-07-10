package config

import (
	"fmt"

	"github.com/caarlos0/env"
	"github.com/spf13/viper"
)

type (
	MainConfig struct {
		Server         ServerConfig
		DatabaseMaster DatabaseMasterConfig
	}

	ServerConfig struct {
		AppName    string `mapstructure:"APP_NAME"    env:"APP_NAME"    envDefault:"go_auth"`
		AppVersion string `mapstructure:"APP_VERSION" env:"APP_VERSION" envDefault:"0.0.1"`
		Host       string `mapstructure:"HOST"        env:"HOST"        envDefault:"0.0.0.0"`
		Port       int    `mapstructure:"PORT"        env:"PORT"        envDefault:"8080"`
	}

	DatabaseMasterConfig struct {
		Host       string `mapstructure:"DB_MASTER_HOST"     env:"DB_MASTER_HOST"`
		Port       int    `mapstructure:"DB_MASTER_PORT"     env:"DB_MASTER_PORT" envDefault:"5432"`
		Dialect    string `mapstructure:"DB_DIALECT"         env:"DB_DIALECT"     envDefault:"postgres"`
		DbName     string `mapstructure:"DB_NAME"            env:"DB_NAME"        envDefault:"simple_crud"`
		DbUser     string `mapstructure:"DB_MASTER_USER"     env:"DB_MASTER_USER"`
		DbPassword string `mapstructure:"DB_MASTER_PASSWORD" env:"DB_MASTER_PASSWORD"`
		TimeZone   string `mapstructure:"DB_TIME_ZONE"       env:"DB_TIME_ZONE" envDefault:"Asia/Jakarta"`
	}
)

func LoadEnv(config *MainConfig) {
	if config == nil {
		panic("config cannot be nil")
	}

	viper.AddConfigPath("./")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	*config = MainConfig{}
	cfgs := []any{
		&config.Server,
		&config.DatabaseMaster,
	}

	if err := viper.ReadInConfig(); err != nil {
		for _, cfgs := range cfgs {
			if err := env.Parse(cfgs); err != nil {
				panic(fmt.Sprintf("cannot parse config from ENV variables, %v", err))
			}
		}
	} else {
		for _, cfgs := range cfgs {
			if err := viper.Unmarshal(cfgs); err != nil {
				panic(fmt.Sprintf("cannot parse config from .env, %v", err))
			}
		}
	}
}
