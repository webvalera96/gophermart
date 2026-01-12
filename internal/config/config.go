// Package config предоставляет функциональность для загрузки и управления конфигурацией приложения.
package config

import (
	"errors"
	"gophermart/internal/logger"
	"os"

	"github.com/spf13/viper"
)

// Config содержит конфигурационные параметры приложения.
// Параметры могут быть установлены через переменные окружения, флаги командной строки или файл .env.
type Config struct {
	RunAddress          string // RunAddress - адрес и порт запуска HTTP сервера (по умолчанию ":8080")
	DatabaseURI         string // DatabaseURI - URI подключения к базе данных PostgreSQL
	AccrualSystemAddress string // AccrualSystemAddress - адрес системы расчёта начислений
}

const configFileName = ".env"

// NewConfig создает новый экземпляр конфигурации.
// Загружает конфигурацию из переменных окружения, флагов командной строки (через env vars) и файла .env.
// Приоритет: переменные окружения > файл .env > значения по умолчанию.
// Принимает logger - логгер для записи информационных сообщений.
// Возвращает указатель на Config с загруженными параметрами.
func NewConfig(logger *logger.Logger) *Config {
	config := &Config{
		RunAddress:          ":8080",
		DatabaseURI:         "",
		AccrualSystemAddress: "",
	}

	// Читаем переменные окружения
	runAddress := os.Getenv("RUN_ADDRESS")
	if runAddress != "" {
		config.RunAddress = runAddress
	}

	databaseURI := os.Getenv("DATABASE_URI")
	if databaseURI != "" {
		config.DatabaseURI = databaseURI
	}

	accrualSystemAddress := os.Getenv("ACCRUAL_SYSTEM_ADDRESS")
	if accrualSystemAddress != "" {
		config.AccrualSystemAddress = accrualSystemAddress
	}

	// Пытаемся прочитать из .env файла (для обратной совместимости)
	var fileLookupError viper.ConfigFileNotFoundError
	viper.SetConfigName(configFileName)
	viper.SetConfigType("env")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		if !errors.As(err, &fileLookupError) {
			logger.Infof("Error reading config file: %v \n", err)
		}
	} else {
		logger.Info("Loaded configuration from .env")
		// Используем значения из .env только если они не были установлены через переменные окружения
		if config.RunAddress == ":8080" {
			if viper.IsSet("RUN_ADDRESS") {
				config.RunAddress = viper.GetString("RUN_ADDRESS")
			}
		}
		if config.DatabaseURI == "" {
			if viper.IsSet("DATABASE_URI") {
				config.DatabaseURI = viper.GetString("DATABASE_URI")
			} else {
				// Обратная совместимость: собираем из отдельных полей
				dbName := viper.GetString("POSTGRES_DB")
				user := viper.GetString("POSTGRES_USER")
				password := viper.GetString("POSTGRES_PASSWORD")
				host := viper.GetString("POSTGRES_HOST")
				port := viper.GetString("POSTGRES_PORT")
				if dbName != "" && user != "" && password != "" && host != "" && port != "" {
					config.DatabaseURI = "postgres://" + user + ":" + password + "@" + host + ":" + port + "/" + dbName + "?sslmode=disable"
				}
			}
		}
		if config.AccrualSystemAddress == "" {
			if viper.IsSet("ACCRUAL_SYSTEM_ADDRESS") {
				config.AccrualSystemAddress = viper.GetString("ACCRUAL_SYSTEM_ADDRESS")
			}
		}
	}

	return config
}
