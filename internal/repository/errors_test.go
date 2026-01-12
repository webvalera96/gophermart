package repository

import (
	"testing"
)

func TestUserAlreadyExistsError_Error(t *testing.T) {
	err := &UserAlreadyExistsError{
		Msg:  "test message",
		Code: 409,
	}
	msg := err.Error()
	if msg == "" {
		t.Error("Error() returned empty string")
	}
}

func TestUserNotFoundError_Error(t *testing.T) {
	err := &UserNotFoundError{
		Msg:  "test message",
		Code: 404,
	}
	msg := err.Error()
	if msg == "" {
		t.Error("Error() returned empty string")
	}
}

func TestWrongCredentialsError_Error(t *testing.T) {
	err := &WrongCredentialsError{
		Msg:  "test message",
		Code: 401,
	}
	msg := err.Error()
	if msg == "" {
		t.Error("Error() returned empty string")
	}
}

func TestOrderNotFoundError_Error(t *testing.T) {
	err := &OrderNotFoundError{
		Msg:  "test message",
		Code: 404,
	}
	msg := err.Error()
	if msg == "" {
		t.Error("Error() returned empty string")
	}
}

func TestOrderAlreadyExistsError_Error(t *testing.T) {
	err := &OrderAlreadyExistsError{
		Msg:  "test message",
		Code: 409,
	}
	msg := err.Error()
	if msg == "" {
		t.Error("Error() returned empty string")
	}
}

func TestInsufficientFundsError_Error(t *testing.T) {
	err := &InsufficientFundsError{
		Msg:  "test message",
		Code: 402,
	}
	msg := err.Error()
	if msg == "" {
		t.Error("Error() returned empty string")
	}
}
