package group

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type GroupBaseStorage struct {
	logger *zap.Logger
	conn   *pgxpool.Pool
}

func NewGroupBaseStorage(
	logger *zap.Logger,
	conn *pgxpool.Pool,
) *GroupBaseStorage {
	return &GroupBaseStorage{
		logger: logger,
		conn:   conn,
	}
}
