package service

import (
	"fmt"
	"learning-platform/users/internal/dto"
	"learning-platform/users/internal/models"
	"learning-platform/users/internal/models/enum"
)

type UserSettingsService struct {
	storage UserSettingsStorage
}

type UserSettingsStorage interface {
	CreateUserSettings(userID int64, language enum.UserLanguage) error
	GetUserSettings(userID int64) (*models.UserSettings, error)
	UpdateUserSettings(userSettings dto.UserSettingsRequest) error
	UpdateUserTheme(userID int64, theme enum.UserTheme) error
}

func NewUserSettingsService(
	storage UserSettingsStorage,
) *UserSettingsService {
	return &UserSettingsService{
		storage: storage,
	}
}

func (us *UserSettingsService) CreateUserSettings(userID int64, language enum.UserLanguage) error {
	err := us.storage.CreateUserSettings(userID, language)
	if err != nil {
		return fmt.Errorf("create user settings: %w", err)
	}

	return nil
}

func (us *UserSettingsService) GetUserSettings(userID int64) (*models.UserSettings, error) {
	userSettings, err := us.storage.GetUserSettings(userID)
	if err != nil {
		return nil, fmt.Errorf("get user settings: %w", err)
	}

	return userSettings, nil
}

func (us *UserSettingsService) UpdateUserSettings(userSettings dto.UserSettingsRequest) (*models.UserSettings, error) {
	err := us.storage.UpdateUserSettings(userSettings)
	if err != nil {
		return nil, fmt.Errorf("update user settings: %w", err)
	}

	resSettings, err := us.storage.GetUserSettings(userSettings.UserID)
	if err != nil {
		return nil, fmt.Errorf("update user settings (get): %w", err)
	}

	return resSettings, nil
}

func (us *UserSettingsService) UpdateUserTheme(userID int64, theme enum.UserTheme) error {
	err := us.storage.UpdateUserTheme(userID, theme)
	if err != nil {
		return fmt.Errorf("update user theme: %w", err)
	}

	return nil
}
