package repository

import "gophermart/internal/models"

type DatabaseRepository interface {
	// --- USER operations ---
	// Create (register) new user in database
	CreateUser(user models.User) error

	// Get user by login from database
	GetUserByLogin(login string) (*models.User, error)

	// Get user balance by login
	GetUserBalance(login string) (*models.User, error)

	// Withdraw balance from user account
	WithdrawBalance(login string, orderNumber string, sum float64) error

	// Get all withdrawals by user login
	GetWithdrawalsByUserLogin(login string) ([]models.Withdrawal, error)

	// --- ORDER operations ---
	// Create new order for user in database
	CreateOrder(order models.Order) (*models.Order, error)

	// Get order by number
	GetOrderByNumber(number string) (*models.Order, error)

	// Get all orders by user login
	GetOrdersByUserLogin(login string) ([]models.Order, error)
}
