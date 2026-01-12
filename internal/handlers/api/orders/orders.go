// Package orders предоставляет HTTP обработчики для работы с заказами.
package orders

import (
	"encoding/json"
	"errors"
	"gophermart/internal/logger"
	"gophermart/internal/models"
	"gophermart/internal/repository"
	"gophermart/internal/service"
	serviceErrors "gophermart/internal/service/errors"
	"io"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
)

// PostOrdersHandler обрабатывает HTTP запросы для создания новых заказов.
type PostOrdersHandler struct {
	logger *logger.Logger
	repo   repository.DatabaseRepository
}

// ServeHTTP обрабатывает HTTP запрос на создание нового заказа.
// Ожидает номер заказа в теле запроса в формате text/plain.
// Требует аутентификации (логин должен быть установлен в заголовке Login через middleware).
// Возможные коды ответа:
//   - 200 — номер заказа уже был загружен этим пользователем;
//   - 202 — новый номер заказа принят в обработку;
//   - 400 — неверный формат запроса;
//   - 401 — пользователь не аутентифицирован;
//   - 409 — номер заказа уже был загружен другим пользователем;
//   - 422 — неверный формат номера заказа;
//   - 500 — внутренняя ошибка сервера.
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

// NewPostOrdersHandler создает новый обработчик создания заказов.
// Принимает logger - логгер для записи сообщений, repo - репозиторий для работы с базой данных.
// Возвращает указатель на PostOrdersHandler.
func NewPostOrdersHandler(logger *logger.Logger, repo repository.DatabaseRepository) *PostOrdersHandler {
	return &PostOrdersHandler{
		logger: logger,
		repo:   repo,
	}
}

// GetOrdersHandler обрабатывает HTTP запросы для получения списка заказов пользователя.
type GetOrdersHandler struct {
	logger *logger.Logger
	repo   repository.DatabaseRepository
}

// ServeHTTP обрабатывает HTTP запрос на получение списка заказов пользователя.
// Требует аутентификации (логин должен быть установлен в заголовке Login через middleware).
// Возвращает JSON массив с информацией о заказах.
// Возможные коды ответа:
//   - 200 — успешная обработка запроса;
//   - 204 — нет данных для ответа;
//   - 401 — пользователь не авторизован;
//   - 500 — внутренняя ошибка сервера.
func (h GetOrdersHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	orders, err := service.GetOrdersByUserLogin(h.repo, r.Header.Get("Login"))
	if err != nil {
		http.Error(w, "Failed to get orders", http.StatusInternalServerError)
		return
	}
	if len(orders) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(orders)
	if err != nil {
		http.Error(w, "Failed to encode orders", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// NewGetOrdersHandler создает новый обработчик получения списка заказов.
// Принимает logger - логгер для записи сообщений, repo - репозиторий для работы с базой данных.
// Возвращает указатель на GetOrdersHandler.
func NewGetOrdersHandler(logger *logger.Logger, repo repository.DatabaseRepository) *GetOrdersHandler {
	return &GetOrdersHandler{
		logger: logger,
		repo:   repo,
	}
}

// GetOrderByNumberHandler обрабатывает HTTP запросы для получения информации о заказе по номеру.
type GetOrderByNumberHandler struct {
	logger *logger.Logger
	repo   repository.DatabaseRepository
}

// OrderResponse представляет информацию о заказе в ответе API.
type OrderResponse struct {
	Order   string   `json:"order"`            // Order - номер заказа
	Status  string   `json:"status"`           // Status - статус заказа
	Accrual *float64 `json:"accrual,omitempty"` // Accrual - сумма начисления (опционально)
}

// ServeHTTP обрабатывает HTTP запрос на получение информации о заказе по номеру.
// Номер заказа передается в URL параметре "number".
// Требует аутентификации (логин должен быть установлен в заголовке Login через middleware).
// Возвращает JSON с информацией о заказе.
// Возможные коды ответа:
//   - 200 — успешная обработка запроса;
//   - 204 — заказ не зарегистрирован в системе расчёта;
//   - 401 — пользователь не авторизован;
//   - 500 — внутренняя ошибка сервера.
func (h GetOrderByNumberHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	number := chi.URLParam(r, "number")
	if number == "" {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	login := r.Header.Get("Login")
	if login == "" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	order, err := service.GetOrderByNumber(h.repo, number, login)
	var onfe *repository.OrderNotFoundError
	if errors.As(err, &onfe) {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	if err != nil {
		http.Error(w, "Failed to get order", http.StatusInternalServerError)
		return
	}

	response := OrderResponse{
		Order:  order.Number,
		Status: order.GetStatus(),
	}

	// Добавляем accrual только если он есть
	if accrual := order.GetAccrual(); accrual != nil {
		response.Accrual = accrual
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

// NewGetOrderByNumberHandler создает новый обработчик получения заказа по номеру.
// Принимает logger - логгер для записи сообщений, repo - репозиторий для работы с базой данных.
// Возвращает указатель на GetOrderByNumberHandler.
func NewGetOrderByNumberHandler(logger *logger.Logger, repo repository.DatabaseRepository) *GetOrderByNumberHandler {
	return &GetOrderByNumberHandler{
		logger: logger,
		repo:   repo,
	}
}
