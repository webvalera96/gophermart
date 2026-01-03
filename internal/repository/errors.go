package repository

import "fmt"

type UserAlreadyExistsError struct {
	Msg  string
	Code int
}

func (uae *UserAlreadyExistsError) Error() string {
	return fmt.Sprintf("code %d: %s", uae.Code, uae.Msg)
}

type UserNotFoundError struct {
	Msg  string
	Code int
}

func (unf *UserNotFoundError) Error() string {
	return fmt.Sprintf("code %d: %s", unf.Code, unf.Msg)
}
