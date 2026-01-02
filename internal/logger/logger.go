package logger

import (
	"go.uber.org/zap"
)

type Logger struct {
	zapLogger zap.SugaredLogger
}

func NewLogger() *Logger {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	sugar := logger.Sugar()

	return &Logger{zapLogger: *sugar}
}

func (logger *Logger) Debug(args ...interface{}) {
	logger.zapLogger.Debug(args)
}

func (logger *Logger) Debugf(template string, args ...interface{}) {
	logger.zapLogger.Debugf(template, args)
}

func (logger *Logger) Info(args ...interface{}) {
	logger.zapLogger.Info(args)
}

func (logger *Logger) Infof(template string, args ...interface{}) {
	logger.zapLogger.Infof(template, args)
}

func (logger *Logger) Error(args ...interface{}) {
	logger.zapLogger.Error(args)
}

func (logger *Logger) Errorf(template string, args ...interface{}) {
	logger.zapLogger.Errorf(template, args)
}

func (logger *Logger) Fatal(args ...interface{}) {
	logger.zapLogger.Fatal(args)
}

func (logger *Logger) Fatalf(template string, args ...interface{}) {
	logger.zapLogger.Fatalf(template, args)
}
