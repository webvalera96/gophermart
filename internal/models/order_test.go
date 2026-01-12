package models

import (
	"testing"
	"time"
)

func TestOrder_SetID(t *testing.T) {
	order := &Order{}
	err := order.SetID(789)
	if err != nil {
		t.Errorf("SetID() returned error: %v", err)
	}
	if order.GetID() != 789 {
		t.Errorf("GetID() = %d, want 789", order.GetID())
	}
}

func TestOrder_GetID(t *testing.T) {
	order := &Order{}
	order.SetID(101)
	if order.GetID() != 101 {
		t.Errorf("GetID() = %d, want 101", order.GetID())
	}
}

func TestOrder_GetStatus(t *testing.T) {
	tests := []struct {
		name     string
		order    Order
		expected string
	}{
		{
			name:     "empty status returns REGISTERED",
			order:    Order{},
			expected: "REGISTERED",
		},
		{
			name:     "PROCESSING status",
			order:    Order{status: "PROCESSING"},
			expected: "PROCESSING",
		},
		{
			name:     "INVALID status",
			order:    Order{status: "INVALID"},
			expected: "INVALID",
		},
		{
			name:     "PROCESSED status",
			order:    Order{status: "PROCESSED"},
			expected: "PROCESSED",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.order.GetStatus(); got != tt.expected {
				t.Errorf("GetStatus() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestOrder_SetStatus(t *testing.T) {
	tests := []struct {
		name    string
		status  string
		wantErr bool
	}{
		{"valid REGISTERED", "REGISTERED", false},
		{"valid PROCESSING", "PROCESSING", false},
		{"valid INVALID", "INVALID", false},
		{"valid PROCESSED", "PROCESSED", false},
		{"invalid status", "INVALID_STATUS", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			order := &Order{}
			err := order.SetStatus(tt.status)
			if (err != nil) != tt.wantErr {
				t.Errorf("SetStatus() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && order.GetStatus() != tt.status {
				t.Errorf("SetStatus() status = %v, want %v", order.GetStatus(), tt.status)
			}
		})
	}
}

func TestOrder_SetAccrual(t *testing.T) {
	order := &Order{}
	accrual := 150.50
	order.SetAccrual(&accrual)
	if order.GetAccrual() == nil || *order.GetAccrual() != accrual {
		t.Errorf("GetAccrual() = %v, want %f", order.GetAccrual(), accrual)
	}
}

func TestOrder_GetAccrual(t *testing.T) {
	order := &Order{}
	if order.GetAccrual() != nil {
		t.Errorf("GetAccrual() = %v, want nil", order.GetAccrual())
	}

	accrual := 200.75
	order.SetAccrual(&accrual)
	if order.GetAccrual() == nil || *order.GetAccrual() != accrual {
		t.Errorf("GetAccrual() = %v, want %f", order.GetAccrual(), accrual)
	}
}

func TestOrder_SetOrderDate(t *testing.T) {
	order := &Order{}
	now := time.Now()
	err := order.SetOrderDate(now)
	if err != nil {
		t.Errorf("SetOrderDate() returned error: %v", err)
	}
	dateStr, err := order.GetOrderDate()
	if err != nil {
		t.Errorf("GetOrderDate() returned error: %v", err)
	}
	parsed, err := time.Parse(time.RFC3339, dateStr)
	if err != nil {
		t.Errorf("Failed to parse date: %v", err)
	}
	if !parsed.Equal(now.Truncate(time.Second)) {
		t.Errorf("GetOrderDate() = %v, want %v", parsed, now.Truncate(time.Second))
	}
}

func TestOrder_GetOrderDate(t *testing.T) {
	order := &Order{}
	now := time.Now()
	order.SetOrderDate(now)
	dateStr, err := order.GetOrderDate()
	if err != nil {
		t.Errorf("GetOrderDate() returned error: %v", err)
	}
	if dateStr == "" {
		t.Error("GetOrderDate() returned empty string")
	}
}
