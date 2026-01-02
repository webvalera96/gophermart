package config

import (
	"errors"
	"gophermart/internal/logger"

	"github.com/spf13/viper"
)

type DatabaseConfig struct {
	DBName   string
	User     string
	Password string
	Host     string
	Port     string
}

type Config struct {
	DatabaseConfig
}

const configFileName = ".env"

func NewConfig(logger *logger.Logger) *Config {

	var fileLookupError viper.ConfigFileNotFoundError

	viper.SetConfigName(configFileName)
	viper.SetConfigType("env")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		if errors.As(err, &fileLookupError) {
			logger.Infof("Config file not found: %v \n", err)
			logger.Info("Loading default configuration")
			return &Config{
				DatabaseConfig: DatabaseConfig{
					DBName:   "postgres",
					User:     "postgres",
					Password: "gophermart",
					Host:     "5432",
					Port:     "localhost",
				},
			}
		}
	}

	logger.Info("Loaded configuration from .env")

	return &Config{
		DatabaseConfig: DatabaseConfig{
			DBName:   viper.GetString("POSTGRES_DB"),
			User:     viper.GetString("POSTGRES_USER"),
			Password: viper.GetString("POSTGRES_PASSWORD"),
			Host:     viper.GetString("POSTGRES_HOST"),
			Port:     viper.GetString("POSTGRES_PORT"),
		},
	}
}
