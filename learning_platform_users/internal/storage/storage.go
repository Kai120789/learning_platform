package storage

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type Storage struct {
	UserStorage         *UserStorage
	UserInfoStorage     *UserInfoStorage
	UserSettingsStorage *UserSettingsStorage
}

func New(
	logger *zap.Logger,
	conn *pgxpool.Pool,
) *Storage {
	return &Storage{
		UserStorage:         NewUserStorage(logger, conn),
		UserInfoStorage:     NewUserInfoStorage(logger, conn),
		UserSettingsStorage: NewUserSettingsStorage(logger, conn),
	}
}

func Connection(connectStr string) (*pgxpool.Pool, error) {
	dbConn, err := pgxpool.New(context.Background(), connectStr)
	if err != nil {
		fmt.Errorf("unable connect to postgres: ", zap.Error(err))
	}
	return dbConn, nil
}
