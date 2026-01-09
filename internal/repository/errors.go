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

type WrongCredentialsError struct {
	Msg  string
	Code int
}

func (wce *WrongCredentialsError) Error() string {
	return fmt.Sprintf("code %d: %s", wce.Code, wce.Msg)
}

type OrderNotFoundError struct {
	Msg  string
	Code int
}

func (onf *OrderNotFoundError) Error() string {
	return fmt.Sprintf("code %d: %s", onf.Code, onf.Msg)
}

type OrderAlreadyExistsError struct {
	Msg  string
	Code int
}

func (oae *OrderAlreadyExistsError) Error() string {
	return fmt.Sprintf("code %d: %s", oae.Code, oae.Msg)
}
