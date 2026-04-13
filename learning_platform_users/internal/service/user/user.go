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
	storage *UserStorage,
) *UserService {
	userInfoService := NewUserInfoService(storage.UserInfoStorage)
	userSettingsService := NewUserSettingsService(storage.UserSettingsStorage)
	userBaseService := NewUserBaseService(storage.UserBaseStorage, userInfoService, userSettingsService)

	return &UserService{
		UserBaseService:     userBaseService,
		UserInfoService:     userInfoService,
		UserSettingsService: userSettingsService,
	}
}
