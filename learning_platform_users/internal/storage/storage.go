package storage

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
	"learning-platform/users/internal/storage/group"
	"learning-platform/users/internal/storage/user"
)

type Storage struct {
	UserStorage  *user.UserStorage
	GroupStorage *group.GroupStorage
}

func New(
	conn *pgxpool.Pool,
) *Storage {
	return &Storage{
		UserStorage:  user.NewUserStorage(conn),
		GroupStorage: group.NewGroupStorage(conn),
	}
}

func Connection(connectStr string) (*pgxpool.Pool, error) {
	dbConn, err := pgxpool.New(context.Background(), connectStr)
	if err != nil {
		fmt.Errorf("unable connect to postgres: ", zap.Error(err))
	}
	return dbConn, nil
}
