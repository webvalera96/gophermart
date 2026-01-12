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

func TestNewLoginHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockDatabaseRepository(ctrl)
	log := logger.NewLogger()

	handler := NewLoginHandler(log, mockRepo)
	if handler == nil {
		t.Error("NewLoginHandler() returned nil")
	}
}

func TestLoginHandler_ServeHTTP_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockDatabaseRepository(ctrl)
	log := logger.NewLogger()
	handler := NewLoginHandler(log, mockRepo)

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

	body, _ := json.Marshal(user)
	req := httptest.NewRequest("POST", "/api/user/login", bytes.NewBuffer(body))
	w := httptest.NewRecorder()

	handler.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("ServeHTTP() status = %d, want %d", w.Code, http.StatusOK)
	}
	if w.Header().Get("Authorization") == "" {
		t.Error("ServeHTTP() Authorization header is empty")
	}
}

func TestLoginHandler_ServeHTTP_WrongCredentials(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockDatabaseRepository(ctrl)
	log := logger.NewLogger()
	handler := NewLoginHandler(log, mockRepo)

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

	body, _ := json.Marshal(user)
	req := httptest.NewRequest("POST", "/api/user/login", bytes.NewBuffer(body))
	w := httptest.NewRecorder()

	handler.ServeHTTP(w, req)

	if w.Code != http.StatusUnauthorized {
		t.Errorf("ServeHTTP() status = %d, want %d", w.Code, http.StatusUnauthorized)
	}
}

func TestLoginHandler_ServeHTTP_BadRequest(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockDatabaseRepository(ctrl)
	log := logger.NewLogger()
	handler := NewLoginHandler(log, mockRepo)

	req := httptest.NewRequest("POST", "/api/user/login", bytes.NewBufferString("invalid json"))
	w := httptest.NewRecorder()

	handler.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("ServeHTTP() status = %d, want %d", w.Code, http.StatusBadRequest)
	}
}
