package models

import (
	"fmt"
	"slices"
	"time"
)

// Order представляет заказ в системе.
// Содержит информацию о номере заказа, статусе, дате загрузки и начислении.
type Order struct {
	id        int
	Login     string   // Login - логин пользователя, которому принадлежит заказ
	Number    string   `json:"number"` // Number - номер заказа
	orderDate time.Time `json:"uploaded_at"` // orderDate - дата загрузки заказа
	status    string    `json:"status"` // status - статус заказа
	accrual   *float64  `json:"accrual,omitempty"` // accrual - сумма начисления (опционально)
}

// GetStatus возвращает статус заказа.
// Если статус не задан, возвращает "REGISTERED" по умолчанию.
// Возвращает строку со статусом заказа.
func (o Order) GetStatus() string {
	if o.status == "" { // если статус не задан, считаем, что заказ зарегистрирован
		return "REGISTERED"
	}
	return o.status
}

// SetStatus устанавливает статус заказа.
// Принимает status - статус заказа. Допустимые значения: "REGISTERED", "PROCESSING", "INVALID", "PROCESSED".
// Возвращает ошибку, если передан недопустимый статус.
func (o *Order) SetStatus(status string) error {
	statuses := []string{"REGISTERED", "PROCESSING", "INVALID", "PROCESSED"}
	if slices.Contains(statuses, status) == false {
		return fmt.Errorf("Invalid status value")
	}

	o.status = status
	return nil
}

// SetAccrual устанавливает сумму начисления для заказа.
// Принимает accrual - указатель на сумму начисления (может быть nil).
func (o *Order) SetAccrual(accrual *float64) {
	o.accrual = accrual
}

// GetAccrual возвращает сумму начисления для заказа.
// Возвращает указатель на сумму начисления или nil, если начисление не установлено.
func (o Order) GetAccrual() *float64 {
	return o.accrual
}

// SetID устанавливает идентификатор заказа.
// Принимает id - идентификатор заказа.
// Возвращает ошибку, если операция не удалась (в текущей реализации всегда nil).
func (o *Order) SetID(id int) error {
	o.id = id
	return nil
}

// GetID возвращает идентификатор заказа.
// Возвращает id - идентификатор заказа.
func (o Order) GetID() int {
	return o.id
}

// GetOrderDate возвращает дату заказа в формате RFC3339.
// Возвращает строку с датой в формате RFC3339 и ошибку, если форматирование не удалось.
func (o Order) GetOrderDate() (string, error) {
	// преобразуем результата в формат RFC3339

	rfc3339 := o.orderDate.Format(time.RFC3339)
	return rfc3339, nil
}

// SetOrderDate устанавливает дату заказа.
// Принимает orderDate - время создания заказа.
// Возвращает ошибку, если операция не удалась (в текущей реализации всегда nil).
func (o *Order) SetOrderDate(orderDate time.Time) error {
	o.orderDate = orderDate
	return nil
}
