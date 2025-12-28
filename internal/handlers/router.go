package handlers

import (
	"gophermart/internal/handlers/api/user"

	"github.com/go-chi/chi/v5"
)

func NewRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Post("/api/user/register", user.RegisterHandler{}.ServeHTTP)
	return r
}
