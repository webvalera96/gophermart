package service

import (
	"gophermart/internal/models"
	"gophermart/internal/repository"
)

func RegisterUser(
	repo repository.DatabaseRepository,
	u models.User) (string, error) {
	err := repo.CreateUser(u)
	if err != nil {
		return "", err
	}
	token, err := IssueToken(u.Login)
	if err != nil {
		return "", err
	}
	return token, nil
}

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

	token, err := IssueToken(u.Login)
	if err != nil {
		return "", err
	}
	return token, nil
}

func GetUserBalance(
	repo repository.DatabaseRepository,
	login string,
) (*models.User, error) {
	user, err := repo.GetUserBalance(login)
	if err != nil {
		return nil, err
	}
	return user, nil
}
