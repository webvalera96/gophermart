package models

import "time"

// Withdrawal представляет операцию списания средств со счета пользователя.
// Содержит информацию о номере заказа, сумме списания и времени обработки.
type Withdrawal struct {
	id          int
	OrderNumber string    `json:"order"` // OrderNumber - номер заказа, по которому произведено списание
	Sum         float64   `json:"sum"` // Sum - сумма списания
	ProcessedAt time.Time `json:"processed_at"` // ProcessedAt - время обработки списания
}

// SetID устанавливает идентификатор операции списания.
// Принимает id - идентификатор операции.
func (w *Withdrawal) SetID(id int) {
	w.id = id
}

// GetID возвращает идентификатор операции списания.
// Возвращает id - идентификатор операции.
func (w Withdrawal) GetID() int {
	return w.id
}

// SetOrderNumber устанавливает номер заказа для операции списания.
// Принимает orderNumber - номер заказа.
func (w *Withdrawal) SetOrderNumber(orderNumber string) {
	w.OrderNumber = orderNumber
}

// GetOrderNumber возвращает номер заказа операции списания.
// Возвращает номер заказа.
func (w Withdrawal) GetOrderNumber() string {
	return w.OrderNumber
}

// SetSum устанавливает сумму списания.
// Принимает sum - сумма списания.
func (w *Withdrawal) SetSum(sum float64) {
	w.Sum = sum
}

// GetSum возвращает сумму списания.
// Возвращает сумму списания.
func (w Withdrawal) GetSum() float64 {
	return w.Sum
}

// SetProcessedAt устанавливает время обработки операции списания.
// Принимает processedAt - время обработки.
func (w *Withdrawal) SetProcessedAt(processedAt time.Time) {
	w.ProcessedAt = processedAt
}

// GetProcessedAt возвращает время обработки операции списания.
// Возвращает время обработки.
func (w Withdrawal) GetProcessedAt() time.Time {
	return w.ProcessedAt
}

// GetProcessedAtString возвращает время обработки операции списания в формате RFC3339.
// Возвращает строку с временем в формате RFC3339.
func (w Withdrawal) GetProcessedAtString() string {
	return w.ProcessedAt.Format(time.RFC3339)
}
