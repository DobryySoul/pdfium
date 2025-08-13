package handlers

import (
	"context"
	"fmt"
	"path/filepath"
	"strings"
	"time"

	"github.com/DobryySoul/PDFium/internal/usecase"
	"github.com/gofiber/fiber/v3"
	"go.uber.org/zap"
)

type PdfHandler struct {
	pdfUC usecase.PdfUC
	l     *zap.Logger
}

func NewPdfHandler(ctx context.Context, pdfUC usecase.PdfUC, l *zap.Logger) (*PdfHandler, error) {
	return &PdfHandler{
		pdfUC: pdfUC,
		l:     l,
	}, nil
}

func (p *PdfHandler) UploadPdfFile(c fiber.Ctx) error {
	file, err := c.FormFile("file")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	filename := extractFilename(file.Header.Get("Content-Disposition"))
	if filename == "" {
		filename = file.Filename
	}

	safeName := generateSafeFilename(filename)

	if err := c.SaveFile(file, "./uploads/"+safeName); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Cannot save file: " + err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status":   "success",
		"filename": safeName,
	})
}

func extractFilename(header string) string {
	parts := strings.SplitSeq(header, ";")
	for part := range parts {
		part = strings.TrimSpace(part)
		if after, ok := strings.CutPrefix(part, "filename="); ok {
			name := after
			return strings.Trim(name, `"`)
		}
		if after, ok := strings.CutPrefix(part, "filename*="); ok {
			name := after
			return strings.Trim(name, `"`)
		}
	}
	return ""
}

func generateSafeFilename(original string) string {
	ext := filepath.Ext(original)
	name := strings.TrimSuffix(original, ext)
	name = strings.ReplaceAll(name, " ", "_")
	return fmt.Sprintf("%d_%s%s", time.Now().UnixNano(), name, ext)
}

// func (p *PdfHandler) ConverterToImage(c fiber.Ctx) error {

// }
