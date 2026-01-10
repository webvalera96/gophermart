package errors

import "fmt"

type OrderAlreadyCreatedByAnotherUser struct {
	Msg string
}

func (oacbau *OrderAlreadyCreatedByAnotherUser) Error() string {
	return fmt.Sprintf("order already created by another user: %s", oacbau.Msg)
}

type OrderAlreadyCreatedByUser struct {
	Msg string
}

func (oacbu *OrderAlreadyCreatedByUser) Error() string {
	return fmt.Sprintf("order already created by user: %s", oacbu.Msg)
}

type OrderInvalidNumberError struct {
	Msg string
}

func (oine *OrderInvalidNumberError) Error() string {
	return fmt.Sprintf("order invalid number error: %s", oine.Msg)
}
