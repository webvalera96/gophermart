// Package server предоставляет функциональность для запуска HTTP сервера.
package server

import (
	"context"
	"fmt"
	"gophermart/internal/config"
	"net"
	"net/http"

	"github.com/go-chi/chi/v5"
	"go.uber.org/fx"
)

// NewHTTPServer создает новый HTTP сервер с указанным роутером и конфигурацией.
// Настраивает lifecycle хуки для запуска и остановки сервера.
// Принимает lc - lifecycle для управления жизненным циклом сервера,
// r - HTTP роутер, cfg - конфигурация приложения с адресом запуска.
// Возвращает настроенный HTTP сервер.
func NewHTTPServer(lc fx.Lifecycle, r *chi.Mux, cfg *config.Config) *http.Server {
	srv := &http.Server{Addr: cfg.RunAddress, Handler: r}
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			ln, err := net.Listen("tcp", srv.Addr)
			if err != nil {
				return err
			}
			fmt.Println("Starting HTTP server at", srv.Addr)
			go srv.Serve(ln)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return srv.Shutdown(ctx)
		},
	})
	return srv
}
