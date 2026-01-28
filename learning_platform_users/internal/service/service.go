package service

import (
	"go.uber.org/zap"
)

type Service struct {
	UserService         *UserService
	UserInfoService     *UserInfoService
	UserSettingsService *UserSettingsService
	RegisterUseCase     *RegisterUseCase
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
	userService := NewUserService(logger, storage.UserStorage)
	userInfoService := NewUserInfoService(logger, storage.UserInfoStorage)
	userSettingsService := NewUserSettingsService(logger, storage.UserSettingsStorage)

	registerStorage := NewRegisterStorageAdapter(
		storage.UserStorage,
		storage.UserInfoStorage,
		storage.UserSettingsStorage,
	)

	return &Service{
		UserService:         userService,
		UserInfoService:     userInfoService,
		UserSettingsService: userSettingsService,
		RegisterUseCase:     NewRegisterUseCase(logger, registerStorage),
	}
}
