package main

import (
	"gophermart/internal/handlers"
	"gophermart/internal/logger"
	"gophermart/internal/repository/pg"
	"gophermart/internal/server"
	"net/http"

	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Provide(server.NewHTTPServer),
		fx.Provide(handlers.NewRouter),
		fx.Provide(pg.NewPGDatabase),
		fx.Provide(logger.NewLogger),
		fx.Invoke(func(*http.Server) {}),
	).Run()
}
