package usecase

import (
	"context"
	"mime/multipart"

	"github.com/DobryySoul/PDFium/internal/repository"
	"go.uber.org/zap"
)

type PdfUC interface {
	SaveUploadFile(ctx context.Context, file *multipart.FileHeader, userID uint) (string, error)
}

type PdfUsecase struct {
	pdfRepo repository.PdfRepo
	logger  *zap.Logger
}

func NewPdfUsecase(ctx context.Context, pdfRepo repository.PdfRepo, logger *zap.Logger) (PdfUC, error) {
	return &PdfUsecase{
		pdfRepo: pdfRepo,
		logger:  logger,
	}, nil
}

func (p *PdfUsecase) SaveUploadFile(ctx context.Context, file *multipart.FileHeader, userID uint) (string, error) {
	return "name_file", nil
}
