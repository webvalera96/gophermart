package user

import (
	"encoding/json"
	"gophermart/internal/logger"
	"gophermart/internal/models"
	"gophermart/internal/repository"
	"net/http"
)

type LoginHandler struct {
	logger *logger.Logger
	repo   repository.DatabaseRepository
}

func NewLoginHandler(
	logger *logger.Logger,
	repo repository.DatabaseRepository,
) LoginHandler {
	return LoginHandler{logger: logger, repo: repo}
}

// 200 — пользователь успешно аутентифицирован;
// 400 — неверный формат запроса;
// 401 — неверная пара логин/пароль;
// 500 — внутренняя ошибка сервера.
func (h LoginHandler) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request,
) {
	var u models.User

	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	user, err := h.repo.GetUserByLogin(u.Login)
	if err != nil {
		http.Error(w, "Wrong login or password", http.StatusUnauthorized)
		return
	}

	if user.Password != u.Password {
		http.Error(w, "Wrong login or password", http.StatusUnauthorized)
		return
	} else {
		h.logger.Info("Successful user login")
		w.WriteHeader(http.StatusOK)
	}
}
