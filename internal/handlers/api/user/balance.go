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

type BalanceHandler struct {
	logger *logger.Logger
	repo   repository.DatabaseRepository
}

type BalanceResponse struct {
	Current   float64 `json:"current"`
	Withdrawn float64 `json:"withdrawn"`
}

func NewBalanceHandler(
	logger *logger.Logger,
	repo repository.DatabaseRepository,
) *BalanceHandler {
	return &BalanceHandler{logger: logger, repo: repo}
}

// 200 — успешная обработка запроса.
// 401 — пользователь не авторизован.
// 500 — внутренняя ошибка сервера.
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

type WithdrawRequest struct {
	Order string  `json:"order"`
	Sum   float64 `json:"sum"`
}

type WithdrawHandler struct {
	logger *logger.Logger
	repo   repository.DatabaseRepository
}

func NewWithdrawHandler(
	logger *logger.Logger,
	repo repository.DatabaseRepository,
) *WithdrawHandler {
	return &WithdrawHandler{logger: logger, repo: repo}
}

// 200 — успешная обработка запроса.
// 401 — пользователь не авторизован.
// 402 — на счету недостаточно средств.
// 422 — неверный номер заказа.
// 500 — внутренняя ошибка сервера.
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
