package config

import (
	"context"
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Host           string         `yaml:"host"`
	Port           string         `yaml:"port"`
	PostgresConfig PostgresConfig `yaml:"postgres"`
	RedisConfig    RedisConfig    `yaml:"redis"`
}

type PostgresConfig struct {
	Host     string `yaml:"host" default:"postgres"`
	Port     string `yaml:"port" default:"5432"`
	Username string `yaml:"username" default:"postgres"`
	Password string `yaml:"password" default:"05042007PULlup!"`
	Database string `yaml:"database" default:"postgres"`
	MaxConns int    `yaml:"max_conn" default:"15"`
	MinConns int    `yaml:"min_conn" default:"0"`
}

type RedisConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}

func LoadConfig(ctx context.Context) (*Config, error) {
	var cfg Config

	if err := cleanenv.ReadConfig("./config/config.yaml", &cfg); err != nil {
		return nil, fmt.Errorf("failed to load config: %w", err)
	}

	return &cfg, nil
}
