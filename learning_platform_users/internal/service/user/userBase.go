package user

import (
	"fmt"
	"github.com/Kai120789/learning_platform_models/models"
	"learning-platform/users/internal/dto"
)

type UserBaseService struct {
	storage             UserBaseStorage
	userInfoService     UserInfo
	userSettingsService UserSettings
}

type UserBaseStorage interface {
	CreateUser(userDto dto.CreateUser) (*int64, error)
	GetUserById(userId int64) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
	ChangePassword(userId int64, newPasswordHash string) error
	ChangeEmail(userId int64, newEmail string) error
}

type UserInfo interface {
	CreateUserInfo(userId int64, userDto dto.CreateUser) error
	GetUserInfo(userId int64) (*models.UserInfo, error)
}

type UserSettings interface {
	CreateUserSettings(userId int64) error
	GetUserSettings(userId int64) (*models.UserSettings, error)
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
	userId, err := s.storage.CreateUser(userDto)
	if err != nil {
		return nil, fmt.Errorf("create user: %w", err)
	}

	err = s.userInfoService.CreateUserInfo(*userId, userDto)
	if err != nil {
		return nil, fmt.Errorf("create user (info): %w", err)
	}

	err = s.userSettingsService.CreateUserSettings(*userId)
	if err != nil {
		return nil, fmt.Errorf("create user (settings): %w", err)
	}

	return userId, nil
}

func (s *UserBaseService) GetUserData(userId int64) (*dto.UserData, error) {
	user, err := s.storage.GetUserById(userId)
	if err != nil {
		return nil, fmt.Errorf("get user data: %w", err)
	}

	userInfo, err := s.userInfoService.GetUserInfo(userId)
	if err != nil {
		return nil, fmt.Errorf("get user data (info): %w", err)
	}

	userSettings, err := s.userSettingsService.GetUserSettings(userId)
	if err != nil {
		return nil, fmt.Errorf("get user data (settings): %w", err)
	}

	return formUserDto(user, userInfo, userSettings), nil
}

func (s *UserBaseService) GetUserById(userId int64) (*models.User, error) {
	user, err := s.storage.GetUserById(userId)
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

func (s *UserBaseService) ChangePassword(userId int64, newPasswordHash string) error {
	err := s.storage.ChangePassword(userId, newPasswordHash)
	if err != nil {
		return fmt.Errorf("change password: %w", err)
	}
	return nil
}
func (s *UserBaseService) ChangeEmail(userId int64, newEmail string) error {
	err := s.storage.ChangeEmail(userId, newEmail)
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
		UserId: user.Id,
		Email:  user.Email,
		UserInfo: dto.UserInfo{
			Name:     userInfo.Name,
			Surname:  userInfo.Surname,
			Lastname: &userInfo.Lastname.String,
			City:     &userInfo.City.String,
			About:    &userInfo.About.String,
			Role:     userInfo.Role,
			Status:   userInfo.Status,
			Class:    &userInfo.Class.Int64,
		},
		UserSettings: dto.UserSettings{
			Is2FaEnabled:           userSettings.Is2FaEnabled,
			IsNotificationsEnabled: userSettings.IsNotificationsEnabled,
		},
	}
}
