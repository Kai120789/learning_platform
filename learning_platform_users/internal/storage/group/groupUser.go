package group

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type GroupUserStorage struct {
	logger *zap.Logger
	conn   *pgxpool.Pool
}

func NewGroupUserStorage(
	logger *zap.Logger,
	conn *pgxpool.Pool,
) *GroupUserStorage {
	return &GroupUserStorage{
		logger: logger,
		conn:   conn,
	}
}
