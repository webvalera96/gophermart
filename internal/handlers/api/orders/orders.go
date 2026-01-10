package orders

import (
	"errors"
	"gophermart/internal/logger"
	"gophermart/internal/models"
	"gophermart/internal/repository"
	"gophermart/internal/service"
	serviceErrors "gophermart/internal/service/errors"
	"io"
	"net/http"
	"strings"
)

type PostOrdersHandler struct {
	logger *logger.Logger
	repo   repository.DatabaseRepository
}

// 200 — номер заказа уже был загружен этим пользователем;
// 202 — новый номер заказа принят в обработку;
// 400 — неверный формат запроса;
// 401 — пользователь не аутентифицирован;
// 409 — номер заказа уже был загружен другим пользователем;
// 422 — неверный формат номера заказа;
// 500 — внутренняя ошибка сервера.
func (h PostOrdersHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}

	if r.Header.Get("Content-Type") != "text/plain" {
		http.Error(w, "Content-Type must be text/plain", http.StatusInternalServerError)
		return
	}

	// TODO: validate that bodyBytes contains a valid order number
	number := strings.Trim(string(bodyBytes), " ")

	// получаем логин аутентифицированного пользователя из заголовков
	login := r.Header.Get("Login")

	savedOrder, err := service.CreateOrder(h.repo, models.Order{Number: number, Login: login})
	var oac *serviceErrors.OrderAlreadyCreatedByUser
	if errors.As(err, &oac) {
		w.WriteHeader(http.StatusOK)
		return
	}
	var oine *serviceErrors.OrderInvalidNumberError
	if errors.As(err, &oine) {
		w.WriteHeader(http.StatusUnprocessableEntity)
	}

	var oaca *serviceErrors.OrderAlreadyCreatedByAnotherUser
	if errors.As(err, &oaca) {
		http.Error(w, "Order created by another user", http.StatusConflict)
	}

	if err != nil {
		http.Error(w, "Failed to create order", http.StatusInternalServerError)
		return
	}

	h.logger.Infof("Order created: %v", savedOrder)
	w.WriteHeader(http.StatusAccepted)
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
