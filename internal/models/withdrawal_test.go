package models

import (
	"testing"
	"time"
)

func TestWithdrawal_SetID(t *testing.T) {
	withdrawal := &Withdrawal{}
	withdrawal.SetID(111)
	if withdrawal.GetID() != 111 {
		t.Errorf("GetID() = %d, want 111", withdrawal.GetID())
	}
}

func TestWithdrawal_GetID(t *testing.T) {
	withdrawal := &Withdrawal{}
	withdrawal.SetID(222)
	if withdrawal.GetID() != 222 {
		t.Errorf("GetID() = %d, want 222", withdrawal.GetID())
	}
}

func TestWithdrawal_SetOrderNumber(t *testing.T) {
	withdrawal := &Withdrawal{}
	orderNumber := "1234567890"
	withdrawal.SetOrderNumber(orderNumber)
	if withdrawal.GetOrderNumber() != orderNumber {
		t.Errorf("GetOrderNumber() = %s, want %s", withdrawal.GetOrderNumber(), orderNumber)
	}
}

func TestWithdrawal_GetOrderNumber(t *testing.T) {
	withdrawal := &Withdrawal{}
	orderNumber := "9876543210"
	withdrawal.SetOrderNumber(orderNumber)
	if withdrawal.GetOrderNumber() != orderNumber {
		t.Errorf("GetOrderNumber() = %s, want %s", withdrawal.GetOrderNumber(), orderNumber)
	}
}

func TestWithdrawal_SetSum(t *testing.T) {
	withdrawal := &Withdrawal{}
	sum := 300.50
	withdrawal.SetSum(sum)
	if withdrawal.GetSum() != sum {
		t.Errorf("GetSum() = %f, want %f", withdrawal.GetSum(), sum)
	}
}

func TestWithdrawal_GetSum(t *testing.T) {
	withdrawal := &Withdrawal{}
	sum := 400.75
	withdrawal.SetSum(sum)
	if withdrawal.GetSum() != sum {
		t.Errorf("GetSum() = %f, want %f", withdrawal.GetSum(), sum)
	}
}

func TestWithdrawal_SetProcessedAt(t *testing.T) {
	withdrawal := &Withdrawal{}
	now := time.Now()
	withdrawal.SetProcessedAt(now)
	if !withdrawal.GetProcessedAt().Equal(now) {
		t.Errorf("GetProcessedAt() = %v, want %v", withdrawal.GetProcessedAt(), now)
	}
}

func TestWithdrawal_GetProcessedAt(t *testing.T) {
	withdrawal := &Withdrawal{}
	now := time.Now()
	withdrawal.SetProcessedAt(now)
	if !withdrawal.GetProcessedAt().Equal(now) {
		t.Errorf("GetProcessedAt() = %v, want %v", withdrawal.GetProcessedAt(), now)
	}
}

func TestWithdrawal_GetProcessedAtString(t *testing.T) {
	withdrawal := &Withdrawal{}
	now := time.Now()
	withdrawal.SetProcessedAt(now)
	str := withdrawal.GetProcessedAtString()
	if str == "" {
		t.Error("GetProcessedAtString() returned empty string")
	}
	parsed, err := time.Parse(time.RFC3339, str)
	if err != nil {
		t.Errorf("Failed to parse date string: %v", err)
	}
	if !parsed.Equal(now.Truncate(time.Second)) {
		t.Errorf("GetProcessedAtString() = %v, want %v", parsed, now.Truncate(time.Second))
	}
}
