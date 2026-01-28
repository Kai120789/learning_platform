package service

import (
	"github.com/Kai120789/learning_platform_models/models"
	"go.uber.org/zap"
	"learning-platform/users/internal/dto"
)

type UserSettingsService struct {
	logger  *zap.Logger
	storage UserSettingsStorage
}

type UserSettingsStorage interface {
	CreateUserSettings(userId int64) error
	GetUserSettings(userId int64) (*models.UserSettings, error)
	UpdateUserSettings(userSettings dto.UserSettings) error
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

func (s *UserSettingsService) CreateUserSettings(userId int64) error {
	return s.storage.CreateUserSettings(userId)
}

func (s *UserSettingsService) GetUserSettings(userId int64) (*models.UserSettings, error) {
	return s.storage.GetUserSettings(userId)
}

func (s *UserSettingsService) UpdateUserSettings(userSettings dto.UserSettings) (*models.UserSettings, error) {
	err := s.storage.UpdateUserSettings(userSettings)
	if err != nil {
		s.logger.Error("update user settings error", zap.Error(err))
	}

	return s.storage.GetUserSettings(userSettings.UserId)
}
