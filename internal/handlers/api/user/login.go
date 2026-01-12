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

// LoginHandler обрабатывает HTTP запросы для аутентификации пользователей.
type LoginHandler struct {
	logger *logger.Logger
	repo   repository.DatabaseRepository
}

// NewLoginHandler создает новый обработчик аутентификации пользователей.
// Принимает logger - логгер для записи сообщений, repo - репозиторий для работы с базой данных.
// Возвращает указатель на LoginHandler.
func NewLoginHandler(
	logger *logger.Logger,
	repo repository.DatabaseRepository,
) *LoginHandler {
	return &LoginHandler{logger: logger, repo: repo}
}

// ServeHTTP обрабатывает HTTP запрос на аутентификацию пользователя.
// Ожидает JSON с полями Login и Password в теле запроса.
// Возвращает JWT токен в заголовке Authorization при успешной аутентификации.
// Возможные коды ответа:
//   - 200 — пользователь успешно аутентифицирован;
//   - 400 — неверный формат запроса;
//   - 401 — неверная пара логин/пароль;
//   - 500 — внутренняя ошибка сервера.
func (h LoginHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var u models.User

	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	var wce *repository.WrongCredentialsError
	token, err := service.LoginUser(h.repo, u)

	if errors.As(err, &wce) {
		http.Error(w, "Wrong login or password", http.StatusUnauthorized)
		return
	} else if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	h.logger.Info("Successful user login")
	w.Header().Set("Authorization", token)
	w.WriteHeader(http.StatusOK)
}
