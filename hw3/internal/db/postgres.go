package db

import (
	"context"
	"errors"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func NewPostgresDB(ctx context.Context) (*pgxpool.Pool, error) {
	err := godotenv.Load("../.env")
	if err != nil {
		return nil, errors.New("error loading .env file")
	}

	connString := os.Getenv("POSTGRES")
	pool, err := pgxpool.New(ctx, connString)
	if err != nil {
		return nil, errors.New("failed to create connection pool")
	}
	return pool, nil
}
