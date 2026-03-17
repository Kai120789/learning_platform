package user

import (
	"go.uber.org/zap"
)

type UserService struct {
	UserBaseService     *UserBaseService
	UserInfoService     *UserInfoService
	UserSettingsService *UserSettingsService
}

type UserStorage struct {
	UserBaseStorage     UserBaseStorage
	UserInfoStorage     UserInfoStorage
	UserSettingsStorage UserSettingsStorage
}

func NewUserService(
	logger *zap.Logger,
	storage *UserStorage,
) *UserService {
	userInfoService := NewUserInfoService(logger, storage.UserInfoStorage)
	userSettingsService := NewUserSettingsService(logger, storage.UserSettingsStorage)
	userBaseService := NewUserBaseService(logger, storage.UserBaseStorage, userInfoService, userSettingsService)

	return &UserService{
		UserBaseService:     userBaseService,
		UserInfoService:     userInfoService,
		UserSettingsService: userSettingsService,
	}
}
