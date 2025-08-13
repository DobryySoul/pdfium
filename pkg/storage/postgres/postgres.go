package postgres

import (
	"context"
	"fmt"

	"github.com/DobryySoul/PDFium/internal/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewConn(ctx context.Context, cfg *config.PostgresConfig) (*pgxpool.Pool, error) {
	conn, err := pgxpool.New(ctx, fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?pool_max_conns=%d&pool_min_conns=%d",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Database,
		cfg.MaxConns,
		cfg.MinConns,
	))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to postgres: %w", err)
	}

	if err := conn.Ping(ctx); err != nil {
		return nil, fmt.Errorf("postgres is unhealthy: %w", err)
	}

	return conn, nil
}
