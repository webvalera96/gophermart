package handlers

import (
	"gophermart/internal/handlers/api/user"
	"gophermart/internal/logger"
	"gophermart/internal/repository"

	"github.com/go-chi/chi/v5"
)

func NewRouter(logger *logger.Logger, repo repository.DatabaseRepository) *chi.Mux {
	r := chi.NewRouter()

	// Handlers creation
	registerHandler := user.NewRegisterHandler(logger, repo)

	// API definition
	r.Post("/api/user/register", registerHandler.ServeHTTP)

	return r
}
