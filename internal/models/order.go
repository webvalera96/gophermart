package models

import (
	"fmt"
	"slices"
	"time"
)

type Order struct {
	id        int
	Login     string
	Number    string
	orderDate time.Time
	status    string
}

func (o Order) GetStatus() string {
	if o.status == "" { // если статус не задан, считаем, что заказ пустой
		return "NEW"
	}
	return o.status
}

func (o *Order) SetStatus(status string) error {
	statuses := []string{"NEW", "PROCESSING", "INVALID", "PROCESSED"}
	if slices.Contains(statuses, status) == false {
		return fmt.Errorf("Invalid status value")
	}

	o.status = status
	return nil
}

func (o *Order) SetID(id int) error {
	o.id = id
	return nil
}

func (o Order) GetID() int {
	return o.id
}

func (o Order) GetOrderDate() (string, error) {
	// преобразуем результата в формат RFC3339

	rfc3339 := o.orderDate.Format(time.RFC3339)
	return rfc3339, nil
}

func (o *Order) SetOrderDate(orderDate time.Time) error {
	o.orderDate = orderDate
	return nil
}
