// Package errors предоставляет специфичные ошибки для сервисного слоя.
package errors

import "fmt"

// OrderAlreadyCreatedByAnotherUser представляет ошибку, когда заказ уже был создан другим пользователем.
type OrderAlreadyCreatedByAnotherUser struct {
	Msg string // Msg - сообщение об ошибке
}

// Error возвращает строковое представление ошибки.
func (oacbau *OrderAlreadyCreatedByAnotherUser) Error() string {
	return fmt.Sprintf("order already created by another user: %s", oacbau.Msg)
}

// OrderAlreadyCreatedByUser представляет ошибку, когда заказ уже был создан этим пользователем.
type OrderAlreadyCreatedByUser struct {
	Msg string // Msg - сообщение об ошибке
}

// Error возвращает строковое представление ошибки.
func (oacbu *OrderAlreadyCreatedByUser) Error() string {
	return fmt.Sprintf("order already created by user: %s", oacbu.Msg)
}

// OrderInvalidNumberError представляет ошибку, когда номер заказа невалиден (не проходит проверку алгоритмом Луна).
type OrderInvalidNumberError struct {
	Msg string // Msg - сообщение об ошибке
}

// Error возвращает строковое представление ошибки.
func (oine *OrderInvalidNumberError) Error() string {
	return fmt.Sprintf("order invalid number error: %s", oine.Msg)
}
