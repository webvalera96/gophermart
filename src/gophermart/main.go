package main

import (
	"gophermart/internal/handlers"
	"gophermart/internal/server"
	"net/http"

	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Provide(server.NewHTTPServer),
		fx.Provide(handlers.NewRouter),
		fx.Invoke(func(*http.Server) {}),
	).Run()
}
