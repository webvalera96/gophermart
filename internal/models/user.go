// Package models предоставляет модели данных для системы gophermart.
package models

// User представляет пользователя системы.
// Содержит информацию о логине, пароле, текущем балансе и сумме списанных средств.
type User struct {
	id             int
	Login          string // Login - логин пользователя
	Password       string // Password - пароль пользователя
	currentBalance float64
	withdrawn      float64
}

// SetID устанавливает идентификатор пользователя.
// Принимает id - идентификатор пользователя.
// Возвращает ошибку, если операция не удалась (в текущей реализации всегда nil).
func (u *User) SetID(id int) error {
	u.id = id

	return nil
}

// GetID возвращает идентификатор пользователя.
// Возвращает id - идентификатор пользователя.
func (u User) GetID() int {

	return u.id
}

// SetCurrentBalance устанавливает текущий баланс пользователя.
// Принимает balance - сумма текущего баланса.
func (u *User) SetCurrentBalance(balance float64) {
	u.currentBalance = balance
}

// GetCurrentBalance возвращает текущий баланс пользователя.
// Возвращает текущий баланс пользователя.
func (u User) GetCurrentBalance() float64 {
	return u.currentBalance
}

// SetWithdrawn устанавливает сумму списанных средств пользователя.
// Принимает withdrawn - сумма списанных средств.
func (u *User) SetWithdrawn(withdrawn float64) {
	u.withdrawn = withdrawn
}

// GetWithdrawn возвращает сумму списанных средств пользователя.
// Возвращает сумму списанных средств.
func (u User) GetWithdrawn() float64 {
	return u.withdrawn
}
