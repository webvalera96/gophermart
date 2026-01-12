package user

import (
	"bytes"
	"encoding/json"
	"gophermart/internal/logger"
	"gophermart/internal/models"
	"gophermart/internal/repository"
	"net/http"
	"net/http/httptest"
	"testing"

	"go.uber.org/mock/gomock"
	"gophermart/internal/mocks"
)

func TestNewRegisterHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockDatabaseRepository(ctrl)
	log := logger.NewLogger()

	handler := NewRegisterHandler(log, mockRepo)
	if handler == nil {
		t.Error("NewRegisterHandler() returned nil")
	}
}

func TestRegisterHandler_ServeHTTP_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockDatabaseRepository(ctrl)
	log := logger.NewLogger()
	handler := NewRegisterHandler(log, mockRepo)

	user := models.User{
		Login:    "testuser",
		Password: "password123",
	}

	mockRepo.EXPECT().
		CreateUser(user).
		Return(nil)

	body, _ := json.Marshal(user)
	req := httptest.NewRequest("POST", "/api/user/register", bytes.NewBuffer(body))
	w := httptest.NewRecorder()

	handler.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("ServeHTTP() status = %d, want %d", w.Code, http.StatusOK)
	}
	if w.Header().Get("Authorization") == "" {
		t.Error("ServeHTTP() Authorization header is empty")
	}
}

func TestRegisterHandler_ServeHTTP_BadRequest(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockDatabaseRepository(ctrl)
	log := logger.NewLogger()
	handler := NewRegisterHandler(log, mockRepo)

	req := httptest.NewRequest("POST", "/api/user/register", bytes.NewBufferString("invalid json"))
	w := httptest.NewRecorder()

	handler.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("ServeHTTP() status = %d, want %d", w.Code, http.StatusBadRequest)
	}
}

func TestRegisterHandler_ServeHTTP_UserAlreadyExists(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockDatabaseRepository(ctrl)
	log := logger.NewLogger()
	handler := NewRegisterHandler(log, mockRepo)

	user := models.User{
		Login:    "testuser",
		Password: "password123",
	}

	mockRepo.EXPECT().
		CreateUser(user).
		Return(&repository.UserAlreadyExistsError{})

	body, _ := json.Marshal(user)
	req := httptest.NewRequest("POST", "/api/user/register", bytes.NewBuffer(body))
	w := httptest.NewRecorder()

	handler.ServeHTTP(w, req)

	if w.Code != http.StatusConflict {
		t.Errorf("ServeHTTP() status = %d, want %d", w.Code, http.StatusConflict)
	}
}
