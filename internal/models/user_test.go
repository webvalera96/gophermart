package models

import (
	"testing"
)

func TestUser_SetID(t *testing.T) {
	user := &User{}
	err := user.SetID(123)
	if err != nil {
		t.Errorf("SetID() returned error: %v", err)
	}
	if user.GetID() != 123 {
		t.Errorf("GetID() = %d, want 123", user.GetID())
	}
}

func TestUser_GetID(t *testing.T) {
	user := &User{}
	user.SetID(456)
	if user.GetID() != 456 {
		t.Errorf("GetID() = %d, want 456", user.GetID())
	}
}

func TestUser_SetCurrentBalance(t *testing.T) {
	user := &User{}
	user.SetCurrentBalance(100.50)
	if user.GetCurrentBalance() != 100.50 {
		t.Errorf("GetCurrentBalance() = %f, want 100.50", user.GetCurrentBalance())
	}
}

func TestUser_GetCurrentBalance(t *testing.T) {
	user := &User{}
	user.SetCurrentBalance(200.75)
	if user.GetCurrentBalance() != 200.75 {
		t.Errorf("GetCurrentBalance() = %f, want 200.75", user.GetCurrentBalance())
	}
}

func TestUser_SetWithdrawn(t *testing.T) {
	user := &User{}
	user.SetWithdrawn(50.25)
	if user.GetWithdrawn() != 50.25 {
		t.Errorf("GetWithdrawn() = %f, want 50.25", user.GetWithdrawn())
	}
}

func TestUser_GetWithdrawn(t *testing.T) {
	user := &User{}
	user.SetWithdrawn(75.00)
	if user.GetWithdrawn() != 75.00 {
		t.Errorf("GetWithdrawn() = %f, want 75.00", user.GetWithdrawn())
	}
}
