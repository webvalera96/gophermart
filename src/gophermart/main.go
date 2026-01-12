package main

import (
	"flag"
	"gophermart/internal/config"
	"gophermart/internal/handlers"
	"gophermart/internal/logger"
	"gophermart/internal/repository/pg"
	"gophermart/internal/server"
	"net/http"
	"os"

	"go.uber.org/fx"
)

func main() {
	// Парсим флаги командной строки
	runAddress := flag.String("a", "", "адрес и порт запуска сервиса")
	databaseURI := flag.String("d", "", "адрес подключения к базе данных")
	accrualSystemAddress := flag.String("r", "", "адрес системы расчёта начислений")
	flag.Parse()

	// Устанавливаем переменные окружения из флагов, если они были переданы
	if *runAddress != "" {
		os.Setenv("RUN_ADDRESS", *runAddress)
	}
	if *databaseURI != "" {
		os.Setenv("DATABASE_URI", *databaseURI)
	}
	if *accrualSystemAddress != "" {
		os.Setenv("ACCRUAL_SYSTEM_ADDRESS", *accrualSystemAddress)
	}

	fx.New(
		fx.Provide(server.NewHTTPServer),
		fx.Provide(handlers.NewRouter),
		fx.Provide(pg.NewPGDatabase),
		fx.Provide(logger.NewLogger),
		fx.Provide(config.NewConfig),
		fx.Invoke(func(*http.Server) {}),
	).Run()
}
