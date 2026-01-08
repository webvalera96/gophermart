package models

type User struct {
	id       int
	Login    string
	Password string
}

func (u *User) SetID(id int) error {
	u.id = id
	// TODO: implement validation logic
	return nil
}

func (u User) GetID() int {
	return u.id
}
