package service

import (
	"errors"
	"gophermart/internal/models"
	"gophermart/internal/repository"
	serviceErrors "gophermart/internal/service/errors"
	"testing"

	"go.uber.org/mock/gomock"
	"gophermart/internal/mocks"
)

func TestRegisterUser_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockDatabaseRepository(ctrl)
	user := models.User{
		Login:    "testuser",
		Password: "password123",
	}

	mockRepo.EXPECT().
		CreateUser(user).
		Return(nil)

	token, err := RegisterUser(mockRepo, user)
	if err != nil {
		t.Fatalf("RegisterUser() returned error: %v", err)
	}
	if token == "" {
		t.Error("RegisterUser() returned empty token")
	}
}

func TestRegisterUser_UserAlreadyExists(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockDatabaseRepository(ctrl)
	user := models.User{
		Login:    "testuser",
		Password: "password123",
	}

	mockRepo.EXPECT().
		CreateUser(user).
		Return(&repository.UserAlreadyExistsError{})

	token, err := RegisterUser(mockRepo, user)
	if err == nil {
		t.Error("RegisterUser() should return error for existing user")
	}
	if token != "" {
		t.Error("RegisterUser() should return empty token on error")
	}
	var expectedErr *repository.UserAlreadyExistsError
	if !errors.As(err, &expectedErr) {
		t.Errorf("RegisterUser() error type = %T, want %T", err, expectedErr)
	}
}

func TestLoginUser_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockDatabaseRepository(ctrl)
	user := models.User{
		Login:    "testuser",
		Password: "password123",
	}
	existingUser := &models.User{
		Login:    "testuser",
		Password: "password123",
	}

	mockRepo.EXPECT().
		GetUserByLogin("testuser").
		Return(existingUser, nil)

	token, err := LoginUser(mockRepo, user)
	if err != nil {
		t.Fatalf("LoginUser() returned error: %v", err)
	}
	if token == "" {
		t.Error("LoginUser() returned empty token")
	}
}

func TestLoginUser_WrongPassword(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockDatabaseRepository(ctrl)
	user := models.User{
		Login:    "testuser",
		Password: "wrongpassword",
	}
	existingUser := &models.User{
		Login:    "testuser",
		Password: "password123",
	}

	mockRepo.EXPECT().
		GetUserByLogin("testuser").
		Return(existingUser, nil)

	token, err := LoginUser(mockRepo, user)
	if err == nil {
		t.Error("LoginUser() should return error for wrong password")
	}
	if token != "" {
		t.Error("LoginUser() should return empty token on error")
	}
	var expectedErr *repository.WrongCredentialsError
	if !errors.As(err, &expectedErr) {
		t.Errorf("LoginUser() error type = %T, want %T", err, expectedErr)
	}
}

func TestLoginUser_UserNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockDatabaseRepository(ctrl)
	user := models.User{
		Login:    "testuser",
		Password: "password123",
	}

	mockRepo.EXPECT().
		GetUserByLogin("testuser").
		Return(nil, &repository.UserNotFoundError{})

	token, err := LoginUser(mockRepo, user)
	if err == nil {
		t.Error("LoginUser() should return error for non-existent user")
	}
	if token != "" {
		t.Error("LoginUser() should return empty token on error")
	}
	var expectedErr *repository.WrongCredentialsError
	if !errors.As(err, &expectedErr) {
		t.Errorf("LoginUser() error type = %T, want %T", err, expectedErr)
	}
}

func TestGetUserBalance_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockDatabaseRepository(ctrl)
	login := "testuser"
	expectedUser := &models.User{
		Login: login,
	}
	expectedUser.SetCurrentBalance(100.50)
	expectedUser.SetWithdrawn(25.00)

	mockRepo.EXPECT().
		GetUserBalance(login).
		Return(expectedUser, nil)

	user, err := GetUserBalance(mockRepo, login)
	if err != nil {
		t.Fatalf("GetUserBalance() returned error: %v", err)
	}
	if user == nil {
		t.Error("GetUserBalance() returned nil user")
	}
	if user.GetCurrentBalance() != 100.50 {
		t.Errorf("GetUserBalance() balance = %f, want 100.50", user.GetCurrentBalance())
	}
}

func TestWithdrawBalance_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockDatabaseRepository(ctrl)
	login := "testuser"
	orderNumber := "4532015112830366"
	sum := 50.00

	mockRepo.EXPECT().
		WithdrawBalance(login, orderNumber, sum).
		Return(nil)

	err := WithdrawBalance(mockRepo, login, orderNumber, sum)
	if err != nil {
		t.Fatalf("WithdrawBalance() returned error: %v", err)
	}
}

func TestWithdrawBalance_InvalidOrderNumber(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockDatabaseRepository(ctrl)
	login := "testuser"
	orderNumber := "1234567890" // invalid Luhn
	sum := 50.00

	err := WithdrawBalance(mockRepo, login, orderNumber, sum)
	if err == nil {
		t.Error("WithdrawBalance() should return error for invalid order number")
	}
	var expectedErr *serviceErrors.OrderInvalidNumberError
	if !errors.As(err, &expectedErr) {
		t.Errorf("WithdrawBalance() error type = %T, want %T", err, expectedErr)
	}
}

func TestWithdrawBalance_InsufficientFunds(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockDatabaseRepository(ctrl)
	login := "testuser"
	orderNumber := "4532015112830366"
	sum := 50.00

	mockRepo.EXPECT().
		WithdrawBalance(login, orderNumber, sum).
		Return(&repository.InsufficientFundsError{})

	err := WithdrawBalance(mockRepo, login, orderNumber, sum)
	if err == nil {
		t.Error("WithdrawBalance() should return error for insufficient funds")
	}
	var expectedErr *repository.InsufficientFundsError
	if !errors.As(err, &expectedErr) {
		t.Errorf("WithdrawBalance() error type = %T, want %T", err, expectedErr)
	}
}

func TestGetWithdrawalsByUserLogin_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockDatabaseRepository(ctrl)
	login := "testuser"
	expectedWithdrawals := []models.Withdrawal{
		{OrderNumber: "123", Sum: 10.00},
		{OrderNumber: "456", Sum: 20.00},
	}

	mockRepo.EXPECT().
		GetWithdrawalsByUserLogin(login).
		Return(expectedWithdrawals, nil)

	withdrawals, err := GetWithdrawalsByUserLogin(mockRepo, login)
	if err != nil {
		t.Fatalf("GetWithdrawalsByUserLogin() returned error: %v", err)
	}
	if len(withdrawals) != len(expectedWithdrawals) {
		t.Errorf("GetWithdrawalsByUserLogin() returned %d withdrawals, want %d", len(withdrawals), len(expectedWithdrawals))
	}
}
