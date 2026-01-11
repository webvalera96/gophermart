package user

import (
	"encoding/json"
	"gophermart/internal/logger"
	"gophermart/internal/repository"
	"gophermart/internal/service"
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
