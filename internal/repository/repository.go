package repository

import "gophermart/internal/models"

type DatabaseRepository interface {
	// --- USER operations ---
	// Create (register) new user in database
	CreateUser(user models.User) error

	// Get user by login from database
	GetUserByLogin(login string) (*models.User, error)

	// --- ORDER operations ---
	// Create new order for user in database
	CreateOrder(order models.Order) (*models.Order, error)
}
