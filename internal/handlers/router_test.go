package handlers

import (
	"gophermart/internal/logger"
	"net/http"
	"net/http/httptest"
	"testing"

	"go.uber.org/mock/gomock"
	"gophermart/internal/mocks"
)

func TestNewRouter(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockDatabaseRepository(ctrl)
	log := logger.NewLogger()

	router := NewRouter(log, mockRepo)
	if router == nil {
		t.Error("NewRouter() returned nil")
	}
}

func TestAuthMiddleware(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockDatabaseRepository(ctrl)
	log := logger.NewLogger()

	middleware := authMiddleware(mockRepo, log)
	if middleware == nil {
		t.Error("authMiddleware() returned nil")
	}

	// Создаем тестовый handler
	handler := middleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	// Тест без токена
	req := httptest.NewRequest("GET", "/test", nil)
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, req)

	if w.Code != http.StatusUnauthorized {
		t.Errorf("authMiddleware() status = %d, want %d", w.Code, http.StatusUnauthorized)
	}
}
