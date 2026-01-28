package service

import "learning-platform/users/internal/dto"

type RegisterStorageAdapter struct {
	user         UserStorage
	userInfo     UserInfoStorage
	userSettings UserSettingsStorage
}

func NewRegisterStorageAdapter(
	user UserStorage,
	userInfo UserInfoStorage,
	userSettings UserSettingsStorage,
) *RegisterStorageAdapter {
	return &RegisterStorageAdapter{
		user:         user,
		userInfo:     userInfo,
		userSettings: userSettings,
	}
}

func (a *RegisterStorageAdapter) CreateUser(dto dto.CreateUser) (*int64, error) {
	return a.user.CreateUser(dto)
}

func (a *RegisterStorageAdapter) CreateUserInfo(id int64, dto dto.CreateUser) error {
	return a.userInfo.CreateUserInfo(id, dto)
}

func (a *RegisterStorageAdapter) CreateUserSettings(id int64) error {
	return a.userSettings.CreateUserSettings(id)
}
