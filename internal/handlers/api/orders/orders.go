package orders

import (
	"gophermart/internal/logger"
	"gophermart/internal/models"
	"gophermart/internal/repository"
	"gophermart/internal/service"
	"io"
	"net/http"
	"strings"
)

type PostOrdersHandler struct {
	logger *logger.Logger
	repo   repository.DatabaseRepository
}

func (h PostOrdersHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}

	// TODO: validate that bodyBytes contains a valid order number
	number := strings.Trim(string(bodyBytes), " ")

	login := r.Header.Get("Login")

	savedOrder, err := service.CreateOrder(h.repo, models.Order{Number: number, Login: login})

	if err != nil {
		http.Error(w, "Failed to create order", http.StatusInternalServerError)
		return
	}

	h.logger.Infof("Order created: %v", savedOrder)
	w.WriteHeader(http.StatusOK)
}

func NewPostOrdersHandler(logger *logger.Logger, repo repository.DatabaseRepository) *PostOrdersHandler {
	return &PostOrdersHandler{
		logger: logger,
		repo:   repo,
	}
}

type GetOrdersHandler struct {
	logger *logger.Logger
	repo   repository.DatabaseRepository
}

func (h GetOrdersHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func NewGetOrdersHandler(logger *logger.Logger, repo repository.DatabaseRepository) *GetOrdersHandler {
	return &GetOrdersHandler{
		logger: logger,
		repo:   repo,
	}
}
