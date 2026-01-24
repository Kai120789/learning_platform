package postgres

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type PostgresStorage struct {
	connection *pgxpool.Pool
	logger     *zap.Logger
}

func New(connection *pgxpool.Pool, logger *zap.Logger) *PostgresStorage {
	return &PostgresStorage{
		connection: connection,
		logger:     logger,
	}
}

func Connection(connectStr string, logger *zap.Logger) (*pgxpool.Pool, error) {
	dbConn, err := pgxpool.New(context.Background(), connectStr)
	if err != nil {
		logger.Error("unable connect to db", zap.Error(err))
		return nil, err
	}

	return dbConn, nil
}
