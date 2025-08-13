package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/DobryySoul/PDFium/internal/entity"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type AuthRepo interface {
	Register(ctx context.Context, user *entity.User) error
	GetUserByEmail(ctx context.Context, email string) (*entity.User, error)
}

type AuthRepository struct {
	pool   *pgxpool.Pool
	logger *zap.Logger
}

func NewAuthRepo(ctx context.Context, pool *pgxpool.Pool, logger *zap.Logger) (AuthRepo, error) {
	return &AuthRepository{
		pool:   pool,
		logger: logger,
	}, nil
}

func (ar *AuthRepository) Register(ctx context.Context, user *entity.User) error {
	query := `
        INSERT INTO users(email, pass_hash, username, created_at, updated_at) 
        VALUES($1, $2, $3, $4, $5);
    `

	_, err := ar.pool.Exec(ctx, query,
		&user.Email,
		&user.PassHash,
		"",
		time.Now(),
		time.Now(),
	)
	if err != nil {
		return fmt.Errorf("failed to create new user: %v", err)
	}

	return nil
}

func (ar *AuthRepository) GetUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	query := `
		SELECT * FROM users WHERE email = $1;
	`
	var user entity.User
	err := ar.pool.QueryRow(ctx, query, email).Scan(
		&user.ID,
		&user.Email,
		&user.PassHash,
		&user.Username,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("It user is not found in database: %w", err)
	}

	return &user, nil
}
