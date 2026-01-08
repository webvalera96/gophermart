package user

// TODO: проверить работу логина
import (
	"encoding/json"
	"errors"
	"gophermart/internal/logger"
	"gophermart/internal/models"
	"gophermart/internal/repository"
	"gophermart/internal/service"
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
