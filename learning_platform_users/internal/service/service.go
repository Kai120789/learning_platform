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
	userInfoService := NewUserInfoService(logger, storage.UserInfoStorage)
	userSettingsService := NewUserSettingsService(logger, storage.UserSettingsStorage)
	userService := NewUserService(logger, storage.UserStorage, userInfoService, userSettingsService)

	return &Service{
		UserService:         userService,
		UserInfoService:     userInfoService,
		UserSettingsService: userSettingsService,
	}
}
