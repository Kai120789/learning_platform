package service

import (
	"fmt"
	"learning-platform/users/internal/dto"
	"learning-platform/users/internal/models"
)

type UserSettingsService struct {
	storage UserSettingsStorage
}

type UserSettingsStorage interface {
	CreateUserSettings(userID int64) error
	GetUserSettings(userID int64) (*models.UserSettings, error)
	UpdateUserSettings(userSettings dto.UserSettings) error
}

func NewUserSettingsService(
	storage UserSettingsStorage,
) *UserSettingsService {
	return &UserSettingsService{
		storage: storage,
	}
}

func (s *UserSettingsService) CreateUserSettings(userID int64) error {
	err := s.storage.CreateUserSettings(userID)
	if err != nil {
		return fmt.Errorf("create user settings: %w", err)
	}

	return nil
}

func (s *UserSettingsService) GetUserSettings(userID int64) (*models.UserSettings, error) {
	userSettings, err := s.storage.GetUserSettings(userID)
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

	resSettings, err := s.storage.GetUserSettings(userSettings.UserID)
	if err != nil {
		return nil, fmt.Errorf("update user settings (get): %w", err)
	}

	return resSettings, nil
}
