package handlers

import (
	"gophermart/internal/handlers/api/orders"
	"gophermart/internal/handlers/api/user"
	"gophermart/internal/logger"
	"gophermart/internal/repository"
	"gophermart/internal/service"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func authMiddleware(_ repository.DatabaseRepository, logger *logger.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token := r.Header.Get("Authorization")
			if token == "" {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			login, err := service.ValidateToken(token)
			if err != nil {
				http.Error(w, "Invalid token", http.StatusUnauthorized)
				return
			}
			r.Header.Set("Login", login)
			next.ServeHTTP(w, r)
		})
	}
}

func NewRouter(
	logger *logger.Logger,
	repo repository.DatabaseRepository,
) *chi.Mux {
	r := chi.NewRouter()

	r.Post("/api/user/register", user.NewRegisterHandler(logger, repo).ServeHTTP)
	r.Post("/api/user/login", user.NewLoginHandler(logger, repo).ServeHTTP)

	r.With(authMiddleware(repo, logger)).Get("/api/user/orders", orders.NewGetOrdersHandler(logger, repo).ServeHTTP)
	r.With(authMiddleware(repo, logger)).Post("/api/user/orders", orders.NewPostOrdersHandler(logger, repo).ServeHTTP)

	return r
}
