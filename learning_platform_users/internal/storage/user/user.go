package user

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type UserStorage struct {
	UserBaseStorage     *UserBaseStorage
	UserInfoStorage     *UserInfoStorage
	UserSettingsStorage *UserSettingsStorage
}

func NewUserStorage(
	logger *zap.Logger,
	conn *pgxpool.Pool,
) *UserStorage {
	return &UserStorage{
		UserBaseStorage:     NewUserBaseStorage(logger, conn),
		UserInfoStorage:     NewUserInfoStorage(logger, conn),
		UserSettingsStorage: NewUserSettingsStorage(logger, conn),
	}
}
