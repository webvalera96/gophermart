package models

type Order struct {
	id        int
	Login     string
	Number    string
	orderDate string
}

func (o Order) SetID(id int) error {
	o.id = id
	return nil
}

func (o Order) GetID() int {
	return o.id
}

func (o Order) GetOrderDate() string {
	return o.orderDate
}

func (o Order) SetOrderDate(orderDate string) error {
	return nil
}
