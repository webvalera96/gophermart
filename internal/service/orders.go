package service

import (
	"gophermart/internal/models"
	"gophermart/internal/repository"
	serviceErrors "gophermart/internal/service/errors"
)

func isValidLuhn(number string) bool {
	// Check if the string contains only numbers
	for _, r := range number {
		if r < '0' || r > '9' {
			return false
		}
	}

	// Luhn algorithm implementation
	sum := 0
	numDigits := len(number)
	for i := numDigits - 1; i >= 0; i-- {
		digit := int(number[i] - '0')

		if (numDigits-i)%2 == 0 {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}
		sum += digit
	}

	return sum%10 == 0
}

func CreateOrder(
	repo repository.DatabaseRepository,
	order models.Order,
) (*models.Order, error) {

	// Validate number before creating order
	valid := isValidLuhn(order.Number)

	if !valid {
		return nil, &serviceErrors.OrderInvalidNumberError{}
	}

	// Check if order already exists
	existOrder, err := repo.GetOrderByNumber(order.Number)
	if err != nil {
		return nil, err
	}

	// if existing order have another user owner
	if existOrder != nil && existOrder.Login != order.Login { // вернуть ошибку, что такой заказ уже создан другим пользователем
		return nil, &serviceErrors.OrderAlreadyCreatedByAnotherUser{}
	} else if existOrder != nil && existOrder.Login == order.Login { // вернуть ошибку, что пользователь уже создавал такой заказ
		return existOrder, &serviceErrors.OrderAlreadyCreatedByUser{}
	}

	savedOrder, err := repo.CreateOrder(order)

	if err != nil {
		return nil, err
	}

	return savedOrder, nil
}

func GetOrdersByUserLogin(
	repo repository.DatabaseRepository,
	login string,
) ([]models.Order, error) {
	orders, err := repo.GetOrdersByUserLogin(login)
	if err != nil {
		return nil, err
	}

	return orders, nil
}
