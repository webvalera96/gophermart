// Package service предоставляет бизнес-логику для работы с пользователями.
package service

import (
	"gophermart/internal/models"
	"gophermart/internal/repository"
	serviceErrors "gophermart/internal/service/errors"
)

// RegisterUser регистрирует нового пользователя в системе и выдает токен аутентификации.
// Принимает repo - репозиторий для работы с базой данных, u - модель пользователя для регистрации.
// Возвращает JWT токен для аутентификации и ошибку, если регистрация не удалась.
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

// LoginUser аутентифицирует пользователя и выдает токен аутентификации.
// Принимает repo - репозиторий для работы с базой данных, u - модель пользователя с логином и паролем.
// Возвращает JWT токен для аутентификации и ошибку, если аутентификация не удалась (неверные учетные данные).
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

// GetUserBalance получает информацию о балансе пользователя.
// Принимает repo - репозиторий для работы с базой данных, login - логин пользователя.
// Возвращает модель пользователя с информацией о балансе и ошибку, если пользователь не найден.
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

// WithdrawBalance списывает средства со счета пользователя.
// Выполняет валидацию номера заказа по алгоритму Луна перед списанием.
// Принимает repo - репозиторий для работы с базой данных,
// login - логин пользователя, orderNumber - номер заказа, sum - сумма списания.
// Возвращает ошибку, если номер заказа невалиден, средств недостаточно или произошла другая ошибка.
func WithdrawBalance(
	repo repository.DatabaseRepository,
	login string,
	orderNumber string,
	sum float64,
) error {
	// Проверяем валидность номера заказа
	if !isValidLuhn(orderNumber) {
		return &serviceErrors.OrderInvalidNumberError{}
	}

	// Списываем средства
	err := repo.WithdrawBalance(login, orderNumber, sum)
	if err != nil {
		return err
	}

	return nil
}

// GetWithdrawalsByUserLogin получает все операции списания для пользователя.
// Принимает repo - репозиторий для работы с базой данных, login - логин пользователя.
// Возвращает список операций списания и ошибку, если произошла ошибка при получении данных.
func GetWithdrawalsByUserLogin(
	repo repository.DatabaseRepository,
	login string,
) ([]models.Withdrawal, error) {
	withdrawals, err := repo.GetWithdrawalsByUserLogin(login)
	if err != nil {
		return nil, err
	}
	return withdrawals, nil
}
