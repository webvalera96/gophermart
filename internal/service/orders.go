package service

import (
	"fmt"
	"gophermart/internal/models"
	"gophermart/internal/repository"
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
	valid := isValidLuhn(order.Number)
	if !valid {
		return nil, fmt.Errorf("Not valid number")
	}
	savedOrder, err := repo.CreateOrder(order)
	if err != nil {
		return nil, err
	}

	return savedOrder, nil
}
