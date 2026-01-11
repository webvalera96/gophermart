package models

import "time"

type Withdrawal struct {
	id          int
	OrderNumber string    `json:"order"`
	Sum         float64   `json:"sum"`
	ProcessedAt time.Time `json:"processed_at"`
}

func (w *Withdrawal) SetID(id int) {
	w.id = id
}

func (w Withdrawal) GetID() int {
	return w.id
}

func (w *Withdrawal) SetOrderNumber(orderNumber string) {
	w.OrderNumber = orderNumber
}

func (w Withdrawal) GetOrderNumber() string {
	return w.OrderNumber
}

func (w *Withdrawal) SetSum(sum float64) {
	w.Sum = sum
}

func (w Withdrawal) GetSum() float64 {
	return w.Sum
}

func (w *Withdrawal) SetProcessedAt(processedAt time.Time) {
	w.ProcessedAt = processedAt
}

func (w Withdrawal) GetProcessedAt() time.Time {
	return w.ProcessedAt
}

func (w Withdrawal) GetProcessedAtString() string {
	return w.ProcessedAt.Format(time.RFC3339)
}
