package logger

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	Logger *zap.Logger
)

func NewLogger(ctx context.Context) (*zap.Logger, error) {
	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.ISO8601TimeEncoder

	consoleEncoder := zapcore.NewJSONEncoder(config)
	consoleWriter := zapcore.AddSync(os.Stdout)

	logsDir := "logs"
	if err := os.MkdirAll(logsDir, 0750); err != nil {
		panic(fmt.Sprintf("Failed to create logs directory: %v", err))
	}

	logFileName := fmt.Sprintf("fb_ads_%s.log", time.Now().Format("2006-01-02"))

	fullLogFilePath := filepath.Join(logsDir, logFileName)
	logFilePath := filepath.Clean(fullLogFilePath)

	file, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
	if err != nil {
		panic(fmt.Sprintf("Failed to open log file: %v", err))
	}

	fileEncoder := zapcore.NewJSONEncoder(config)
	fileWriter := zapcore.AddSync(file)

	core := zapcore.NewTee(
		zapcore.NewCore(consoleEncoder, consoleWriter, zapcore.InfoLevel),
		zapcore.NewCore(fileEncoder, fileWriter, zapcore.InfoLevel),
	)

	Logger = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))

	Logger.Info("Logger initialized successfully")

	return Logger, nil
}
