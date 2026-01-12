package orders

import (
	"bytes"
	"gophermart/internal/logger"
	"gophermart/internal/models"
	"gophermart/internal/repository"
	"net/http"
	"net/http/httptest"
	"testing"

	"go.uber.org/mock/gomock"
	"gophermart/internal/mocks"
)

func TestNewPostOrdersHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockDatabaseRepository(ctrl)
	log := logger.NewLogger()

	handler := NewPostOrdersHandler(log, mockRepo)
	if handler == nil {
		t.Error("NewPostOrdersHandler() returned nil")
	}
}

func TestPostOrdersHandler_ServeHTTP_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockDatabaseRepository(ctrl)
	log := logger.NewLogger()
	handler := NewPostOrdersHandler(log, mockRepo)

	orderNumber := "4532015112830366"
	order := models.Order{
		Number: orderNumber,
		Login:  "testuser",
	}

	mockRepo.EXPECT().
		GetOrderByNumber(orderNumber).
		Return(nil, &repository.OrderNotFoundError{})

	mockRepo.EXPECT().
		CreateOrder(order).
		Return(&order, nil)

	req := httptest.NewRequest("POST", "/api/user/orders", bytes.NewBufferString(orderNumber))
	req.Header.Set("Content-Type", "text/plain")
	req.Header.Set("Login", "testuser")
	w := httptest.NewRecorder()

	handler.ServeHTTP(w, req)

	if w.Code != http.StatusAccepted {
		t.Errorf("ServeHTTP() status = %d, want %d", w.Code, http.StatusAccepted)
	}
}

func TestNewGetOrdersHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockDatabaseRepository(ctrl)
	log := logger.NewLogger()

	handler := NewGetOrdersHandler(log, mockRepo)
	if handler == nil {
		t.Error("NewGetOrdersHandler() returned nil")
	}
}

func TestGetOrdersHandler_ServeHTTP_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockDatabaseRepository(ctrl)
	log := logger.NewLogger()
	handler := NewGetOrdersHandler(log, mockRepo)

	login := "testuser"
	orders := []models.Order{
		{Number: "123", Login: login},
	}

	mockRepo.EXPECT().
		GetOrdersByUserLogin(login).
		Return(orders, nil)

	req := httptest.NewRequest("GET", "/api/user/orders", nil)
	req.Header.Set("Login", login)
	w := httptest.NewRecorder()

	handler.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("ServeHTTP() status = %d, want %d", w.Code, http.StatusOK)
	}
}

func TestNewGetOrderByNumberHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockDatabaseRepository(ctrl)
	log := logger.NewLogger()

	handler := NewGetOrderByNumberHandler(log, mockRepo)
	if handler == nil {
		t.Error("NewGetOrderByNumberHandler() returned nil")
	}
}
