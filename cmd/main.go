package main

import (
	"context"

	"github.com/DobryySoul/PDFium/internal/app"
	"github.com/DobryySoul/PDFium/pkg/logger"
	"go.uber.org/zap"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	logger, err := logger.NewLogger(ctx)
	if err != nil {
		logger.Error("failed to init logger", zap.Error(err))
		return
	}

	if err := app.Run(ctx, logger); err != nil {
		logger.Fatal("failed run app", zap.Error(err))
	}
}
