package storage

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type UserSettingsStorage struct {
	logger *zap.Logger
	conn   *pgxpool.Pool
}

func NewUserSettingsStorage(
	logger *zap.Logger,
	conn *pgxpool.Pool,
) *UserSettingsStorage {
	return &UserSettingsStorage{
		logger: logger,
		conn:   conn,
	}
}
