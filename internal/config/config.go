package config

import (
	"errors"
	"gophermart/internal/logger"
	"os"

	"github.com/joho/godotenv"
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

func NewConfig(logger logger.Logger) *Config {
	_, err := os.Stat(".env")

	// Если файлов конфъгурации нет, то
	if !errors.Is(err, os.ErrNotExist) {

		logger.Info("Loaded default configuration")

		return &Config{
			DatabaseConfig: DatabaseConfig{
				DBName:   "postgres",
				User:     "postgres",
				Password: "gophermart",
				Host:     "5432",
				Port:     "localhost",
			},
		}
	} else { // получить настройки из файлов переменных окружений
		err := godotenv.Load()
		if err != nil {
			logger.Errorf("Error loading .env file: %v", err)
		}

		logger.Info("Loaded configuration from .env")

		return &Config{
			DatabaseConfig: DatabaseConfig{
				DBName:   os.Getenv("POSTGRES_DB"),
				User:     os.Getenv("POSTGRES_USER"),
				Password: os.Getenv("POSTGRES_PASSWORD"),
				Host:     os.Getenv("POSTGRES_HOST"),
				Port:     os.Getenv("POSTGRES_PORT"),
			},
		}
	}

}
