package service

import (
	"gophermart/internal/models"
	"gophermart/internal/repository"
)

func CreateOrder(
	repo repository.DatabaseRepository,
	order models.Order,
) (*models.Order, error) {
	savedOrder, err := repo.CreateOrder(order)
	if err != nil {
		return nil, err
	}

	return savedOrder, nil
}
