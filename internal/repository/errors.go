package repository

import "fmt"

type UserAlreadyExistsError struct {
	Msg  string
	Code int
}

func (uae *UserAlreadyExistsError) Error() string {
	return fmt.Sprintf("code %d: %s", uae.Code, uae.Msg)
}
