package user

import (
	"encoding/json"
	"errors"
	"gophermart/internal/logger"
	"gophermart/internal/repository"
	"gophermart/internal/service"
	serviceErrors "gophermart/internal/service/errors"
	"net/http"
)

// BalanceHandler обрабатывает HTTP запросы для получения баланса пользователя.
type BalanceHandler struct {
	logger *logger.Logger
	repo   repository.DatabaseRepository
}

// BalanceResponse представляет ответ с информацией о балансе пользователя.
type BalanceResponse struct {
	Current   float64 `json:"current"`   // Current - текущий баланс пользователя
	Withdrawn float64 `json:"withdrawn"` // Withdrawn - сумма списанных средств
}

// NewBalanceHandler создает новый обработчик получения баланса пользователя.
// Принимает logger - логгер для записи сообщений, repo - репозиторий для работы с базой данных.
// Возвращает указатель на BalanceHandler.
func NewBalanceHandler(
	logger *logger.Logger,
	repo repository.DatabaseRepository,
) *BalanceHandler {
	return &BalanceHandler{logger: logger, repo: repo}
}

// ServeHTTP обрабатывает HTTP запрос на получение баланса пользователя.
// Требует аутентификации (логин должен быть установлен в заголовке Login через middleware).
// Возвращает JSON с информацией о балансе.
// Возможные коды ответа:
//   - 200 — успешная обработка запроса;
//   - 401 — пользователь не авторизован;
//   - 500 — внутренняя ошибка сервера.
func (h *BalanceHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	login := r.Header.Get("Login")
	if login == "" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	user, err := service.GetUserBalance(h.repo, login)
	if err != nil {
		http.Error(w, "Failed to get user balance", http.StatusInternalServerError)
		return
	}

	response := BalanceResponse{
		Current:   user.GetCurrentBalance(),
		Withdrawn: user.GetWithdrawn(),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

// WithdrawRequest представляет запрос на списание средств.
type WithdrawRequest struct {
	Order string  `json:"order"` // Order - номер заказа
	Sum   float64 `json:"sum"`   // Sum - сумма списания
}

// WithdrawHandler обрабатывает HTTP запросы для списания средств со счета пользователя.
type WithdrawHandler struct {
	logger *logger.Logger
	repo   repository.DatabaseRepository
}

// NewWithdrawHandler создает новый обработчик списания средств.
// Принимает logger - логгер для записи сообщений, repo - репозиторий для работы с базой данных.
// Возвращает указатель на WithdrawHandler.
func NewWithdrawHandler(
	logger *logger.Logger,
	repo repository.DatabaseRepository,
) *WithdrawHandler {
	return &WithdrawHandler{logger: logger, repo: repo}
}

// ServeHTTP обрабатывает HTTP запрос на списание средств со счета пользователя.
// Ожидает JSON с полями Order и Sum в теле запроса.
// Требует аутентификации (логин должен быть установлен в заголовке Login через middleware).
// Возможные коды ответа:
//   - 200 — успешная обработка запроса;
//   - 401 — пользователь не авторизован;
//   - 402 — на счету недостаточно средств;
//   - 422 — неверный номер заказа;
//   - 500 — внутренняя ошибка сервера.
func (h *WithdrawHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	login := r.Header.Get("Login")
	if login == "" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	var req WithdrawRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err := service.WithdrawBalance(h.repo, login, req.Order, req.Sum)
	var ife *repository.InsufficientFundsError
	if errors.As(err, &ife) {
		http.Error(w, "Insufficient funds", http.StatusPaymentRequired)
		return
	}
	var oine *serviceErrors.OrderInvalidNumberError
	if errors.As(err, &oine) {
		http.Error(w, "Invalid order number", http.StatusUnprocessableEntity)
		return
	}
	if err != nil {
		http.Error(w, "Failed to withdraw balance", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// WithdrawalResponse представляет информацию об операции списания средств.
type WithdrawalResponse struct {
	Order       string  `json:"order"`        // Order - номер заказа
	Sum         float64 `json:"sum"`         // Sum - сумма списания
	ProcessedAt string  `json:"processed_at"` // ProcessedAt - время обработки в формате RFC3339
}

// GetWithdrawalsHandler обрабатывает HTTP запросы для получения списка операций списания пользователя.
type GetWithdrawalsHandler struct {
	logger *logger.Logger
	repo   repository.DatabaseRepository
}

// NewGetWithdrawalsHandler создает новый обработчик получения списка операций списания.
// Принимает logger - логгер для записи сообщений, repo - репозиторий для работы с базой данных.
// Возвращает указатель на GetWithdrawalsHandler.
func NewGetWithdrawalsHandler(
	logger *logger.Logger,
	repo repository.DatabaseRepository,
) *GetWithdrawalsHandler {
	return &GetWithdrawalsHandler{logger: logger, repo: repo}
}

// ServeHTTP обрабатывает HTTP запрос на получение списка операций списания пользователя.
// Требует аутентификации (логин должен быть установлен в заголовке Login через middleware).
// Возвращает JSON массив с информацией об операциях списания.
// Возможные коды ответа:
//   - 200 — успешная обработка запроса;
//   - 204 — нет ни одного списания;
//   - 401 — пользователь не авторизован;
//   - 500 — внутренняя ошибка сервера.
func (h *GetWithdrawalsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	login := r.Header.Get("Login")
	if login == "" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	withdrawals, err := service.GetWithdrawalsByUserLogin(h.repo, login)
	if err != nil {
		http.Error(w, "Failed to get withdrawals", http.StatusInternalServerError)
		return
	}

	if len(withdrawals) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	// Преобразуем withdrawals в формат ответа
	response := make([]WithdrawalResponse, 0, len(withdrawals))
	for _, withdrawal := range withdrawals {
		response = append(response, WithdrawalResponse{
			Order:       withdrawal.GetOrderNumber(),
			Sum:         withdrawal.GetSum(),
			ProcessedAt: withdrawal.GetProcessedAtString(),
		})
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}
