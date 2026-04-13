package user

import (
	"fmt"
	"github.com/Kai120789/learning_platform_models/models"
	"learning-platform/users/internal/dto"
)

type UserSettingsService struct {
	storage UserSettingsStorage
}

type UserSettingsStorage interface {
	CreateUserSettings(userId int64) error
	GetUserSettings(userId int64) (*models.UserSettings, error)
	UpdateUserSettings(userSettings dto.UserSettings) error
}

func NewUserSettingsService(
	storage UserSettingsStorage,
) *UserSettingsService {
	return &UserSettingsService{
		storage: storage,
	}
}

func (s *UserSettingsService) CreateUserSettings(userId int64) error {
	err := s.storage.CreateUserSettings(userId)
	if err != nil {
		return fmt.Errorf("create user settings: %w", err)
	}

	return nil
}

func (s *UserSettingsService) GetUserSettings(userId int64) (*models.UserSettings, error) {
	userSettings, err := s.storage.GetUserSettings(userId)
	if err != nil {
		return nil, fmt.Errorf("get user settings: %w", err)
	}

	return userSettings, nil
}

func (s *UserSettingsService) UpdateUserSettings(userSettings dto.UserSettings) (*models.UserSettings, error) {
	err := s.storage.UpdateUserSettings(userSettings)
	if err != nil {
		return nil, fmt.Errorf("update user settings: %w", err)
	}

	resSettings, err := s.storage.GetUserSettings(userSettings.UserId)
	if err != nil {
		return nil, fmt.Errorf("update user settings (get): %w", err)
	}

	return resSettings, nil
}
