package orders

import (
	"gophermart/internal/logger"
	"gophermart/internal/repository"
	"net/http"
)

type PostOrdersHandler struct {
	logger *logger.Logger
	repo   repository.DatabaseRepository
}

func (h PostOrdersHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

type GetOrdersHandler struct {
	logger *logger.Logger
	repo   repository.DatabaseRepository
}

func (h GetOrdersHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
