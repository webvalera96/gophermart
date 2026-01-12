package user

import (
	"bytes"
	"encoding/json"
	"gophermart/internal/logger"
	"gophermart/internal/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"go.uber.org/mock/gomock"
	"gophermart/internal/mocks"
)

func TestNewBalanceHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockDatabaseRepository(ctrl)
	log := logger.NewLogger()

	handler := NewBalanceHandler(log, mockRepo)
	if handler == nil {
		t.Error("NewBalanceHandler() returned nil")
	}
}

func TestBalanceHandler_ServeHTTP_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockDatabaseRepository(ctrl)
	log := logger.NewLogger()
	handler := NewBalanceHandler(log, mockRepo)

	login := "testuser"
	user := &models.User{Login: login}
	user.SetCurrentBalance(100.50)
	user.SetWithdrawn(25.00)

	mockRepo.EXPECT().
		GetUserBalance(login).
		Return(user, nil)

	req := httptest.NewRequest("GET", "/api/user/balance", nil)
	req.Header.Set("Login", login)
	w := httptest.NewRecorder()

	handler.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("ServeHTTP() status = %d, want %d", w.Code, http.StatusOK)
	}

	var response BalanceResponse
	if err := json.NewDecoder(w.Body).Decode(&response); err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}
	if response.Current != 100.50 {
		t.Errorf("BalanceResponse.Current = %f, want 100.50", response.Current)
	}
}

func TestNewWithdrawHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockDatabaseRepository(ctrl)
	log := logger.NewLogger()

	handler := NewWithdrawHandler(log, mockRepo)
	if handler == nil {
		t.Error("NewWithdrawHandler() returned nil")
	}
}

func TestWithdrawHandler_ServeHTTP_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockDatabaseRepository(ctrl)
	log := logger.NewLogger()
	handler := NewWithdrawHandler(log, mockRepo)

	login := "testuser"
	reqBody := WithdrawRequest{
		Order: "4532015112830366",
		Sum:   50.00,
	}

	mockRepo.EXPECT().
		WithdrawBalance(login, reqBody.Order, reqBody.Sum).
		Return(nil)

	body, _ := json.Marshal(reqBody)
	req := httptest.NewRequest("POST", "/api/user/balance/withdraw", bytes.NewBuffer(body))
	req.Header.Set("Login", login)
	w := httptest.NewRecorder()

	handler.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("ServeHTTP() status = %d, want %d", w.Code, http.StatusOK)
	}
}

func TestNewGetWithdrawalsHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockDatabaseRepository(ctrl)
	log := logger.NewLogger()

	handler := NewGetWithdrawalsHandler(log, mockRepo)
	if handler == nil {
		t.Error("NewGetWithdrawalsHandler() returned nil")
	}
}

func TestGetWithdrawalsHandler_ServeHTTP_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockDatabaseRepository(ctrl)
	log := logger.NewLogger()
	handler := NewGetWithdrawalsHandler(log, mockRepo)

	login := "testuser"
	withdrawals := []models.Withdrawal{
		{OrderNumber: "123", Sum: 10.00},
	}

	mockRepo.EXPECT().
		GetWithdrawalsByUserLogin(login).
		Return(withdrawals, nil)

	req := httptest.NewRequest("GET", "/api/user/withdrawals", nil)
	req.Header.Set("Login", login)
	w := httptest.NewRecorder()

	handler.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("ServeHTTP() status = %d, want %d", w.Code, http.StatusOK)
	}
}

func TestGetWithdrawalsHandler_ServeHTTP_NoContent(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockDatabaseRepository(ctrl)
	log := logger.NewLogger()
	handler := NewGetWithdrawalsHandler(log, mockRepo)

	login := "testuser"

	mockRepo.EXPECT().
		GetWithdrawalsByUserLogin(login).
		Return([]models.Withdrawal{}, nil)

	req := httptest.NewRequest("GET", "/api/user/withdrawals", nil)
	req.Header.Set("Login", login)
	w := httptest.NewRecorder()

	handler.ServeHTTP(w, req)

	if w.Code != http.StatusNoContent {
		t.Errorf("ServeHTTP() status = %d, want %d", w.Code, http.StatusNoContent)
	}
}
