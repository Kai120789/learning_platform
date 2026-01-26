package storage

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type UserInfoStorage struct {
	logger *zap.Logger
	conn   *pgxpool.Pool
}

func NewUserInfoStorage(
	logger *zap.Logger,
	conn *pgxpool.Pool,
) *UserInfoStorage {
	return &UserInfoStorage{
		logger: logger,
		conn:   conn,
	}
}
