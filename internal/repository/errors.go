// Package repository предоставляет специфичные ошибки для слоя репозитория.
package repository

import "fmt"

// UserAlreadyExistsError представляет ошибку, когда пользователь с таким логином уже существует.
type UserAlreadyExistsError struct {
	Msg  string // Msg - сообщение об ошибке
	Code int    // Code - код ошибки
}

// Error возвращает строковое представление ошибки.
func (uae *UserAlreadyExistsError) Error() string {
	return fmt.Sprintf("code %d: %s", uae.Code, uae.Msg)
}

// UserNotFoundError представляет ошибку, когда пользователь не найден в базе данных.
type UserNotFoundError struct {
	Msg  string // Msg - сообщение об ошибке
	Code int    // Code - код ошибки
}

// Error возвращает строковое представление ошибки.
func (unf *UserNotFoundError) Error() string {
	return fmt.Sprintf("code %d: %s", unf.Code, unf.Msg)
}

// WrongCredentialsError представляет ошибку, когда предоставлены неверные учетные данные (логин/пароль).
type WrongCredentialsError struct {
	Msg  string // Msg - сообщение об ошибке
	Code int    // Code - код ошибки
}

// Error возвращает строковое представление ошибки.
func (wce *WrongCredentialsError) Error() string {
	return fmt.Sprintf("code %d: %s", wce.Code, wce.Msg)
}

// OrderNotFoundError представляет ошибку, когда заказ не найден в базе данных.
type OrderNotFoundError struct {
	Msg  string // Msg - сообщение об ошибке
	Code int    // Code - код ошибки
}

// Error возвращает строковое представление ошибки.
func (onf *OrderNotFoundError) Error() string {
	return fmt.Sprintf("code %d: %s", onf.Code, onf.Msg)
}

// OrderAlreadyExistsError представляет ошибку, когда заказ уже существует в базе данных.
type OrderAlreadyExistsError struct {
	Msg  string // Msg - сообщение об ошибке
	Code int    // Code - код ошибки
}

// Error возвращает строковое представление ошибки.
func (oae *OrderAlreadyExistsError) Error() string {
	return fmt.Sprintf("code %d: %s", oae.Code, oae.Msg)
}

// InsufficientFundsError представляет ошибку, когда на счету пользователя недостаточно средств для операции.
type InsufficientFundsError struct {
	Msg  string // Msg - сообщение об ошибке
	Code int    // Code - код ошибки
}

// Error возвращает строковое представление ошибки.
func (ife *InsufficientFundsError) Error() string {
	return fmt.Sprintf("code %d: %s", ife.Code, ife.Msg)
}
