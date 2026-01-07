package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func NewPostgresDB(ctx context.Context) (*pgxpool.Pool, error) {
	err := godotenv.Load("../.env")
	if err != nil {
		return nil, fmt.Errorf("error loading .env file: %w", err)
	}

	connString := os.Getenv("POSTGRES")
	pool, err := pgxpool.New(ctx, connString)
	if err != nil {
		return nil, fmt.Errorf("failed to create connection pool: %w", err)
	}
	return pool, nil
}
