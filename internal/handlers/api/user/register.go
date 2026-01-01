package user

import (
	"encoding/json"
	"errors"
	"gophermart/internal/logger"
	"gophermart/internal/models"
	"gophermart/internal/repository"
	"net/http"
)

type RegisterHandler struct {
	logger *logger.Logger
	repo   repository.DatabaseRepository
}

func NewRegisterHandler(logger *logger.Logger, repo repository.DatabaseRepository) RegisterHandler {
	return RegisterHandler{logger: logger, repo: repo}
}

// 200 http.StatusOK — пользователь успешно зарегистрирован и аутентифицирован;
// 400 http.StatusBadRequest — неверный формат запроса;
// 409 http.StatusConflict — логин уже занят;
// 500 http.StatusInternalServerError — внутренняя ошибка сервера.
func (h *RegisterHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var u models.User

	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err = h.repo.CreateUser(u)

	var uaeError *repository.UserAlreadyExistsError
	if errors.As(err, &uaeError) {
		http.Error(w, "User already exists", http.StatusConflict)
		return
	} else if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	h.logger.Info("Successfuly create user")
	w.WriteHeader(http.StatusOK)
}
