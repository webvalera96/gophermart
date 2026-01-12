package config

import (
	"gophermart/internal/logger"
	"os"
	"testing"
)

func TestNewConfig_DefaultValues(t *testing.T) {
	log := logger.NewLogger()
	cfg := NewConfig(log)

	if cfg.RunAddress != ":8080" {
		t.Errorf("NewConfig() RunAddress = %s, want :8080", cfg.RunAddress)
	}
	if cfg.DatabaseURI != "" {
		t.Errorf("NewConfig() DatabaseURI = %s, want empty", cfg.DatabaseURI)
	}
	if cfg.AccrualSystemAddress != "" {
		t.Errorf("NewConfig() AccrualSystemAddress = %s, want empty", cfg.AccrualSystemAddress)
	}
}

func TestNewConfig_EnvironmentVariables(t *testing.T) {
	// Сохраняем оригинальные значения
	originalRunAddress := os.Getenv("RUN_ADDRESS")
	originalDatabaseURI := os.Getenv("DATABASE_URI")
	originalAccrualSystemAddress := os.Getenv("ACCRUAL_SYSTEM_ADDRESS")

	// Устанавливаем тестовые значения
	os.Setenv("RUN_ADDRESS", ":9090")
	os.Setenv("DATABASE_URI", "postgres://user:pass@localhost:5432/db")
	os.Setenv("ACCRUAL_SYSTEM_ADDRESS", "http://localhost:8081")

	// Восстанавливаем после теста
	defer func() {
		if originalRunAddress != "" {
			os.Setenv("RUN_ADDRESS", originalRunAddress)
		} else {
			os.Unsetenv("RUN_ADDRESS")
		}
		if originalDatabaseURI != "" {
			os.Setenv("DATABASE_URI", originalDatabaseURI)
		} else {
			os.Unsetenv("DATABASE_URI")
		}
		if originalAccrualSystemAddress != "" {
			os.Setenv("ACCRUAL_SYSTEM_ADDRESS", originalAccrualSystemAddress)
		} else {
			os.Unsetenv("ACCRUAL_SYSTEM_ADDRESS")
		}
	}()

	log := logger.NewLogger()
	cfg := NewConfig(log)

	if cfg.RunAddress != ":9090" {
		t.Errorf("NewConfig() RunAddress = %s, want :9090", cfg.RunAddress)
	}
	if cfg.DatabaseURI != "postgres://user:pass@localhost:5432/db" {
		t.Errorf("NewConfig() DatabaseURI = %s, want postgres://user:pass@localhost:5432/db", cfg.DatabaseURI)
	}
	if cfg.AccrualSystemAddress != "http://localhost:8081" {
		t.Errorf("NewConfig() AccrualSystemAddress = %s, want http://localhost:8081", cfg.AccrualSystemAddress)
	}
}
