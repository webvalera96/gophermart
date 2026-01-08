package service

import (
	"gophermart/internal/models"
	"gophermart/internal/repository"
)

// TODO: реализовать генерацию и возврат токена при логине
func RegisterUser(
	// TODO: генерировать авторизационный токен
	repo repository.DatabaseRepository,
	u models.User) (string, error) {
	err := repo.CreateUser(u)
	if err != nil {
		return "", err
	}
	return "", nil
}

// TODO: вместо пустой строки возвращать заголовок Authorization с токеном
func LoginUser(
	repo repository.DatabaseRepository,
	u models.User,
) (string, error) {
	user, err := repo.GetUserByLogin(u.Login)
	// TODO: определять типы ошибок для 500 internal server error
	if err != nil {
		return "", &repository.WrongCredentialsError{}
	}

	if user.Password != u.Password {
		return "", &repository.WrongCredentialsError{}
	}
	return "", nil
}
