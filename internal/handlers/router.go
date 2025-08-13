package handlers

import (
	"context"

	"github.com/gofiber/fiber/v3"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

func NewRouter(ctx context.Context, f *fiber.App, pdf *PdfHandler, auth *AuthHandler, rdb *redis.Client, l *zap.Logger) {

	api := f.Group("/api/v1")

	api.Post("/register", auth.Register)
	api.Post("/login", auth.Login)

	api.Post("/upload", pdf.UploadPdfFile)
	// api.Post("/pdf-converter", pdf.ConverterToImage)
}
