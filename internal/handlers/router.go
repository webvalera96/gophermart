package handlers

import (
	"gophermart/internal/handlers/api/user"
	"gophermart/internal/logger"
	"gophermart/internal/repository"

	"github.com/go-chi/chi/v5"
)

func NewRouter(
	logger *logger.Logger,
	repo repository.DatabaseRepository,
) *chi.Mux {
	r := chi.NewRouter()

	// Создание обработчиков
	registerHandler := user.NewRegisterHandler(logger, repo)
	loginHandler := user.NewLoginHandler(logger, repo)

	// Middleware
	// TODO: реализовать проверку токена в заголовке Authorization

	// Определение API
	r.Post("/api/user/register", registerHandler.ServeHTTP)
	r.Post("/api/user/login", loginHandler.ServeHTTP)

	return r
}
