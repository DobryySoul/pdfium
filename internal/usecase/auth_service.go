package usecase

import (
	"context"
	"fmt"

	"github.com/DobryySoul/PDFium/internal/entity"
	"github.com/DobryySoul/PDFium/internal/repository"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

const (
	customCost = 15
)

type AuthUC interface {
	Register(ctx context.Context, user *entity.DoRegister) error
	Login(ctx context.Context, req *entity.DoLogin) (string, error)
}

type AuthUsecase struct {
	authRepo repository.AuthRepo
	logger   *zap.Logger
}

func NewAuthUsecase(ctx context.Context, authRepo repository.AuthRepo, logger *zap.Logger) (AuthUC, error) {
	return &AuthUsecase{
		authRepo: authRepo,
		logger:   logger,
	}, nil
}

func (au *AuthUsecase) Register(ctx context.Context, userCreate *entity.DoRegister) error {
	passHash, err := bcrypt.GenerateFromPassword([]byte(userCreate.Password), customCost)
	if err != nil {
		return fmt.Errorf("failed to generate password hash: %w", err)
	}

	user := &entity.User{
		Email:    userCreate.Email,
		PassHash: string(passHash),
	}

	if err := au.authRepo.Register(ctx, user); err != nil {
		return fmt.Errorf("failed to create new user: %w", err)
	}

	return nil
}

func (au *AuthUsecase) Login(ctx context.Context, req *entity.DoLogin) (string, error) {
	user, err := au.authRepo.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return "", fmt.Errorf("user not found by email: %w", err)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PassHash), []byte(req.Password)); err != nil {
		return "", fmt.Errorf("invalid password: %w", err)
	}

	return "secret-token", nil
}
