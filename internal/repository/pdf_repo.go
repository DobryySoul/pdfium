package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type PdfRepo interface {
}

type PdfRepository struct {
	pool   *pgxpool.Pool
	logger *zap.Logger
}

func NewPdfRepo(ctx context.Context, pool *pgxpool.Pool, logger *zap.Logger) (PdfRepo, error) {
	return &PdfRepository{
		pool:   pool,
		logger: logger,
	}, nil
}

// func (p *PdfRepository) SaveUploadFile(ctx context.Context, file *multipart.FileHeader, userID uint) (string, error) {
// 	query := `

// 	`

// 	return "", nil
// }
