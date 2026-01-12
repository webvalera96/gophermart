// Package logger предоставляет обёртку над zap логгером для единообразного логирования в приложении.
package logger

import (
	"go.uber.org/zap"
)

// Logger представляет логгер приложения.
// Использует zap.SugaredLogger для удобного логирования с поддержкой форматирования.
type Logger struct {
	zapLogger zap.SugaredLogger
}

// NewLogger создает новый экземпляр логгера с production конфигурацией.
// Возвращает указатель на Logger, готовый к использованию.
func NewLogger() *Logger {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	sugar := logger.Sugar()

	return &Logger{zapLogger: *sugar}
}

// Debug записывает debug сообщение в лог.
// Принимает args - аргументы для логирования.
func (logger *Logger) Debug(args ...interface{}) {
	logger.zapLogger.Debug(args)
}

// Debugf записывает форматированное debug сообщение в лог.
// Принимает template - шаблон сообщения и args - аргументы для подстановки.
func (logger *Logger) Debugf(template string, args ...interface{}) {
	logger.zapLogger.Debugf(template, args)
}

// Info записывает информационное сообщение в лог.
// Принимает args - аргументы для логирования.
func (logger *Logger) Info(args ...interface{}) {
	logger.zapLogger.Info(args)
}

// Infof записывает форматированное информационное сообщение в лог.
// Принимает template - шаблон сообщения и args - аргументы для подстановки.
func (logger *Logger) Infof(template string, args ...interface{}) {
	logger.zapLogger.Infof(template, args)
}

// Error записывает сообщение об ошибке в лог.
// Принимает args - аргументы для логирования.
func (logger *Logger) Error(args ...interface{}) {
	logger.zapLogger.Error(args)
}

// Errorf записывает форматированное сообщение об ошибке в лог.
// Принимает template - шаблон сообщения и args - аргументы для подстановки.
func (logger *Logger) Errorf(template string, args ...interface{}) {
	logger.zapLogger.Errorf(template, args)
}

// Fatal записывает критическое сообщение в лог и завершает выполнение программы.
// Принимает args - аргументы для логирования.
func (logger *Logger) Fatal(args ...interface{}) {
	logger.zapLogger.Fatal(args)
}

// Fatalf записывает форматированное критическое сообщение в лог и завершает выполнение программы.
// Принимает template - шаблон сообщения и args - аргументы для подстановки.
func (logger *Logger) Fatalf(template string, args ...interface{}) {
	logger.zapLogger.Fatalf(template, args)
}
