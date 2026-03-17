package group

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type GroupStorage struct {
	GroupBaseStorage *GroupBaseStorage
	GroupUserStorage *GroupUserStorage
}

func NewGroupStorage(
	logger *zap.Logger,
	conn *pgxpool.Pool,
) *GroupStorage {
	return &GroupStorage{
		GroupUserStorage: NewGroupUserStorage(logger, conn),
		GroupBaseStorage: NewGroupBaseStorage(logger, conn),
	}
}
