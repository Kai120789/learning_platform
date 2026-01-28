package service

import "go.uber.org/zap"

type UserSettingsService struct {
	logger  *zap.Logger
	storage UserSettingsStorage
}

type UserSettingsStorage interface {
	CreateUserSettings(userId int64) error
}

func NewUserSettingsService(
	logger *zap.Logger,
	storage UserSettingsStorage,
) *UserSettingsService {
	return &UserSettingsService{
		logger:  logger,
		storage: storage,
	}
}
