package repository

import "gophermart/internal/models"

type DatabaseRepository interface {
	// Create (register) new user in database
	CreateUser(user models.User) error
}
