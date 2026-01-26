package service

import (
	"go.uber.org/zap"
)

type Service struct {
	UserService         *UserService
	UserInfoService     *UserInfoService
	UserSettingsService *UserSettingsService
}

type Storage struct {
	UserStorage         UserStorage
	UserInfoStorage     UserInfoStorage
	UserSettingsStorage UserSettingsStorage
}

func New(
	logger *zap.Logger,
	storage *Storage,
) *Service {
	return &Service{
		UserService:         NewUserService(logger, &storage.UserStorage),
		UserInfoService:     NewUserInfoService(logger, &storage.UserInfoStorage),
		UserSettingsService: NewUserSettingsService(logger, &storage.UserSettingsStorage),
	}
}
