package app

import (
	"context"
	"fmt"

	"github.com/DobryySoul/PDFium/internal/config"
	"github.com/DobryySoul/PDFium/internal/handlers"
	"github.com/DobryySoul/PDFium/internal/repository"
	"github.com/DobryySoul/PDFium/internal/usecase"
	"github.com/DobryySoul/PDFium/pkg/redis"
	"github.com/DobryySoul/PDFium/pkg/storage/postgres"
	"github.com/gofiber/fiber/v3"
	"go.uber.org/zap"
)

func Run(ctx context.Context, l *zap.Logger) error {
	cfg, err := config.LoadConfig(ctx)
	if err != nil {
		return fmt.Errorf("failed to load initial config: %w", err)
	}

	pgpool, err := postgres.NewConn(ctx, &cfg.PostgresConfig)
	if err != nil {
		return fmt.Errorf("failed to init connection pool with postgres: %w", err)
	}
	rdb, err := redis.NewRedisClient(ctx, &cfg.RedisConfig)
	if err != nil {
		return fmt.Errorf("failed to init redis client: %w", err)
	}

	pdfRepo, err := repository.NewPdfRepo(ctx, pgpool, l)
	if err != nil {
		return fmt.Errorf("failed to create pdf repository: %w", err)
	}
	authRepo, err := repository.NewAuthRepo(ctx, pgpool, l)
	if err != nil {
		return fmt.Errorf("failed to create auth repository: %w", err)
	}

	pdfUC, err := usecase.NewPdfUsecase(ctx, pdfRepo, l)
	if err != nil {
		return fmt.Errorf("failed to create pdf usecase: %w", err)
	}
	authfUC, err := usecase.NewAuthUsecase(ctx, authRepo, l)
	if err != nil {
		return fmt.Errorf("failed to create auth usecase: %w", err)
	}

	pdfHandler, err := handlers.NewPdfHandler(ctx, pdfUC, l)
	if err != nil {
		return fmt.Errorf("failed to create pdf handler: %w", err)
	}
	authHandler, err := handlers.NewAuthHandler(ctx, authfUC, l)
	if err != nil {
		return fmt.Errorf("failed to create auth handler: %w", err)
	}

	f := fiber.New()

	handlers.NewRouter(ctx, f, pdfHandler, authHandler, rdb, l)

	f.Listen(fmt.Sprintf("%s:%s", cfg.Host, cfg.Port))
	return nil
}
