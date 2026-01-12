// Package user предоставляет HTTP обработчики для работы с пользователями.
package user

import (
	"encoding/json"
	"errors"
	"gophermart/internal/logger"
	"gophermart/internal/models"
	"gophermart/internal/repository"
	"gophermart/internal/service"
	"net/http"
)

// RegisterHandler обрабатывает HTTP запросы для регистрации новых пользователей.
type RegisterHandler struct {
	logger *logger.Logger
	repo   repository.DatabaseRepository
}

// NewRegisterHandler создает новый обработчик регистрации пользователей.
// Принимает logger - логгер для записи сообщений, repo - репозиторий для работы с базой данных.
// Возвращает указатель на RegisterHandler.
func NewRegisterHandler(
	logger *logger.Logger,
	repo repository.DatabaseRepository,
) *RegisterHandler {
	return &RegisterHandler{logger: logger, repo: repo}
}

// ServeHTTP обрабатывает HTTP запрос на регистрацию пользователя.
// Ожидает JSON с полями Login и Password в теле запроса.
// Возвращает JWT токен в заголовке Authorization при успешной регистрации.
// Возможные коды ответа:
//   - 200 http.StatusOK — пользователь успешно зарегистрирован и аутентифицирован;
//   - 400 http.StatusBadRequest — неверный формат запроса;
//   - 409 http.StatusConflict — логин уже занят;
//   - 500 http.StatusInternalServerError — внутренняя ошибка сервера.
func (h *RegisterHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var u models.User

	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	var uaeError *repository.UserAlreadyExistsError
	token, err := service.RegisterUser(h.repo, u)
	// err = h.repo.CreateUser(u)

	if errors.As(err, &uaeError) {
		http.Error(w, "User already exists", http.StatusConflict)
		return
	} else if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	h.logger.Info("Successfuly create user")
	w.Header().Set("Authorization", token)
	w.WriteHeader(http.StatusOK)
}
