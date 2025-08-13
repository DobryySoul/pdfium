package handlers

import (
	"context"

	"github.com/DobryySoul/PDFium/internal/entity"
	"github.com/DobryySoul/PDFium/internal/usecase"
	"github.com/gofiber/fiber/v3"
	"go.uber.org/zap"
)

type AuthHandler struct {
	authUC usecase.AuthUC
	l      *zap.Logger
}

func NewAuthHandler(ctx context.Context, authUC usecase.AuthUC, l *zap.Logger) (*AuthHandler, error) {
	return &AuthHandler{
		authUC: authUC,
		l:      l,
	}, nil
}

func (au *AuthHandler) Register(c fiber.Ctx) error {
	var req *entity.DoRegister
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if err := au.authUC.Register(c.Context(), req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	c.Status(fiber.StatusCreated)
	return nil
}

func (au *AuthHandler) Login(c fiber.Ctx) error {
	var req *entity.DoLogin
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	token, err := au.authUC.Login(c.Context(), req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"token": token,
	})
}
