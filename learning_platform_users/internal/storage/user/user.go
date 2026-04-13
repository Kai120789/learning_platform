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
	conn *pgxpool.Pool,
) *UserStorage {
	return &UserStorage{
		UserBaseStorage:     NewUserBaseStorage(conn),
		UserInfoStorage:     NewUserInfoStorage(conn),
		UserSettingsStorage: NewUserSettingsStorage(conn),
	}
}
