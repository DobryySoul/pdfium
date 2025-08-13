package redis

import (
	"context"
	"fmt"

	"github.com/DobryySoul/PDFium/internal/config"

	"github.com/redis/go-redis/v9"
)

func NewRedisClient(ctx context.Context, cfg *config.RedisConfig) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
		Username: cfg.Username,
		Password: cfg.Password,
		DB:       cfg.DB,
	})

	// if err := rdb.Ping(ctx); err != nil {
	// 	return nil, fmt.Errorf("failed to check healthy redis: %v", err)
	// }

	return rdb, nil
}
