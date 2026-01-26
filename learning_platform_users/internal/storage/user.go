package storage

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type UserStorage struct {
	logger *zap.Logger
	conn   *pgxpool.Pool
}

func NewUserStorage(
	logger *zap.Logger,
	conn *pgxpool.Pool,
) *UserStorage {
	return &UserStorage{
		logger: logger,
		conn:   conn,
	}
}
