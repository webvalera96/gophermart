package logger

import (
	"testing"
)

func TestNewLogger(t *testing.T) {
	logger := NewLogger()
	if logger == nil {
		t.Error("NewLogger() returned nil")
	}
}

func TestLogger_Debug(t *testing.T) {
	logger := NewLogger()
	// Просто проверяем, что метод не паникует
	logger.Debug("test message")
}

func TestLogger_Debugf(t *testing.T) {
	logger := NewLogger()
	// Просто проверяем, что метод не паникует
	logger.Debugf("test message: %s", "value")
}

func TestLogger_Info(t *testing.T) {
	logger := NewLogger()
	// Просто проверяем, что метод не паникует
	logger.Info("test message")
}

func TestLogger_Infof(t *testing.T) {
	logger := NewLogger()
	// Просто проверяем, что метод не паникует
	logger.Infof("test message: %s", "value")
}

func TestLogger_Error(t *testing.T) {
	logger := NewLogger()
	// Просто проверяем, что метод не паникует
	logger.Error("test error")
}

func TestLogger_Errorf(t *testing.T) {
	logger := NewLogger()
	// Просто проверяем, что метод не паникует
	logger.Errorf("test error: %s", "value")
}
