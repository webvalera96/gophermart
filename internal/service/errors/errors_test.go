package errors

import (
	"testing"
)

func TestOrderAlreadyCreatedByAnotherUser_Error(t *testing.T) {
	err := &OrderAlreadyCreatedByAnotherUser{
		Msg: "test message",
	}
	msg := err.Error()
	if msg == "" {
		t.Error("Error() returned empty string")
	}
}

func TestOrderAlreadyCreatedByUser_Error(t *testing.T) {
	err := &OrderAlreadyCreatedByUser{
		Msg: "test message",
	}
	msg := err.Error()
	if msg == "" {
		t.Error("Error() returned empty string")
	}
}

func TestOrderInvalidNumberError_Error(t *testing.T) {
	err := &OrderInvalidNumberError{
		Msg: "test message",
	}
	msg := err.Error()
	if msg == "" {
		t.Error("Error() returned empty string")
	}
}
