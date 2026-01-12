// Package repository определяет интерфейсы для работы с хранилищем данных.
package repository

import "gophermart/internal/models"

// DatabaseRepository определяет интерфейс для работы с базой данных.
// Предоставляет методы для управления пользователями, заказами и операциями списания средств.
type DatabaseRepository interface {
	// --- USER operations ---
	// CreateUser создает (регистрирует) нового пользователя в базе данных.
	// Принимает user - модель пользователя для создания.
	// Возвращает ошибку, если пользователь уже существует или произошла другая ошибка.
	CreateUser(user models.User) error

	// GetUserByLogin получает пользователя по логину из базы данных.
	// Принимает login - логин пользователя.
	// Возвращает указатель на модель пользователя и ошибку, если пользователь не найден.
	GetUserByLogin(login string) (*models.User, error)

	// GetUserBalance получает информацию о балансе пользователя по логину.
	// Принимает login - логин пользователя.
	// Возвращает указатель на модель пользователя с информацией о балансе и ошибку, если пользователь не найден.
	GetUserBalance(login string) (*models.User, error)

	// WithdrawBalance списывает средства со счета пользователя.
	// Принимает login - логин пользователя, orderNumber - номер заказа, sum - сумма списания.
	// Возвращает ошибку, если средств недостаточно или произошла другая ошибка.
	WithdrawBalance(login string, orderNumber string, sum float64) error

	// GetWithdrawalsByUserLogin получает все операции списания для пользователя по логину.
	// Принимает login - логин пользователя.
	// Возвращает список операций списания и ошибку, если произошла ошибка при получении данных.
	GetWithdrawalsByUserLogin(login string) ([]models.Withdrawal, error)

	// --- ORDER operations ---
	// CreateOrder создает новый заказ для пользователя в базе данных.
	// Принимает order - модель заказа для создания.
	// Возвращает указатель на созданный заказ и ошибку, если произошла ошибка.
	CreateOrder(order models.Order) (*models.Order, error)

	// GetOrderByNumber получает заказ по номеру из базы данных.
	// Принимает number - номер заказа.
	// Возвращает указатель на модель заказа и ошибку, если заказ не найден.
	GetOrderByNumber(number string) (*models.Order, error)

	// GetOrdersByUserLogin получает все заказы пользователя по логину.
	// Принимает login - логин пользователя.
	// Возвращает список заказов и ошибку, если произошла ошибка при получении данных.
	GetOrdersByUserLogin(login string) ([]models.Order, error)
}
