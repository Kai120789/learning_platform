package service

import (
	"fmt"
	"learning-platform/users/internal/dto"
	"learning-platform/users/internal/models"
	"learning-platform/users/internal/utils"
)

type UserBaseService struct {
	storage             UserBaseStorage
	userInfoService     UserInfo
	userSettingsService UserSettings
}

type UserBaseStorage interface {
	CreateUser(userDto dto.CreateUser) (*int64, error)
	GetUserById(userID int64) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
	ChangePassword(userID int64, newPasswordHash string) error
	ChangeEmail(userID int64, newEmail string) error
}

type UserInfo interface {
	CreateUserInfo(userID int64, userDto dto.CreateUser) error
	GetUserInfo(userID int64) (*models.UserInfo, error)
}

type UserSettings interface {
	CreateUserSettings(userID int64) error
	GetUserSettings(userID int64) (*models.UserSettings, error)
}

func NewUserBaseService(
	storage UserBaseStorage,
	userInfoService UserInfo,
	userSettingsService UserSettings,
) *UserBaseService {
	return &UserBaseService{
		storage:             storage,
		userInfoService:     userInfoService,
		userSettingsService: userSettingsService,
	}
}

func (s *UserBaseService) CreateUser(userDto dto.CreateUser) (*int64, error) {
	userID, err := s.storage.CreateUser(userDto)
	if err != nil {
		return nil, fmt.Errorf("create user: %w", err)
	}

	err = s.userInfoService.CreateUserInfo(*userID, userDto)
	if err != nil {
		return nil, fmt.Errorf("create user (info): %w", err)
	}

	err = s.userSettingsService.CreateUserSettings(*userID)
	if err != nil {
		return nil, fmt.Errorf("create user (settings): %w", err)
	}

	return userID, nil
}

func (s *UserBaseService) GetUserData(userID int64) (*dto.UserData, error) {
	user, err := s.storage.GetUserById(userID)
	if err != nil {
		return nil, fmt.Errorf("get user data: %w", err)
	}

	userInfo, err := s.userInfoService.GetUserInfo(userID)
	if err != nil {
		return nil, fmt.Errorf("get user data (info): %w", err)
	}

	userSettings, err := s.userSettingsService.GetUserSettings(userID)
	if err != nil {
		return nil, fmt.Errorf("get user data (settings): %w", err)
	}

	return formUserDto(user, userInfo, userSettings), nil
}

func (s *UserBaseService) GetUserById(userID int64) (*models.User, error) {
	user, err := s.storage.GetUserById(userID)
	if err != nil {
		return nil, fmt.Errorf("get user by id: %w", err)
	}

	return user, nil
}

func (s *UserBaseService) GetUserByEmail(email string) (*models.User, error) {
	user, err := s.storage.GetUserByEmail(email)
	if err != nil {
		return nil, fmt.Errorf("get user by email: %w", err)
	}
	return user, nil
}

func (s *UserBaseService) ChangePassword(userID int64, newPasswordHash string) error {
	err := s.storage.ChangePassword(userID, newPasswordHash)
	if err != nil {
		return fmt.Errorf("change password: %w", err)
	}
	return nil
}
func (s *UserBaseService) ChangeEmail(userID int64, newEmail string) error {
	err := s.storage.ChangeEmail(userID, newEmail)
	if err != nil {
		return fmt.Errorf("change email: %w", err)
	}
	return nil
}

func formUserDto(
	user *models.User,
	userInfo *models.UserInfo,
	userSettings *models.UserSettings,
) *dto.UserData {
	return &dto.UserData{
		UserID: user.ID,
		Email:  user.Email,
		Role:   user.Role,
		Status: user.Status,
		UserInfo: dto.UserInfo{
			Name:     userInfo.Name,
			Surname:  userInfo.Surname,
			Lastname: utils.DBStringToOptional(userInfo.Lastname),
			City:     utils.DBStringToOptional(userInfo.City),
			About:    utils.DBStringToOptional(userInfo.About),
		},
		UserSettings: dto.UserSettings{
			Is2FaEnabled:           userSettings.Is2FaEnabled,
			IsNotificationsEnabled: userSettings.IsNotificationsEnabled,
		},
	}
}
