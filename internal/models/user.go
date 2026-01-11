package models

type User struct {
	id             int
	Login          string
	Password       string
	currentBalance float64
	withdrawn      float64
}

func (u *User) SetID(id int) error {
	u.id = id

	return nil
}

func (u User) GetID() int {

	return u.id
}

func (u *User) SetCurrentBalance(balance float64) {
	u.currentBalance = balance
}

func (u User) GetCurrentBalance() float64 {
	return u.currentBalance
}

func (u *User) SetWithdrawn(withdrawn float64) {
	u.withdrawn = withdrawn
}

func (u User) GetWithdrawn() float64 {
	return u.withdrawn
}
