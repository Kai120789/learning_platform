package service

import (
	"github.com/Kai120789/learning_platform_models/models"
	"go.uber.org/zap"
	"learning-platform/users/internal/dto"
)

type UserService struct {
	logger              *zap.Logger
	storage             UserStorage
	userInfoService     UserInfo
	userSettingsService UserSettings
}

type UserStorage interface {
	CreateUser(userDto dto.CreateUser) (*int64, error)
	GetUser(userId int64) (*models.User, error)
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

func NewUserService(
	logger *zap.Logger,
	storage UserStorage,
	userInfoService UserInfo,
	userSettingsService UserSettings,
) *UserService {
	return &UserService{
		logger:              logger,
		storage:             storage,
		userInfoService:     userInfoService,
		userSettingsService: userSettingsService,
	}
}

func (s *UserService) CreateUser(userDto dto.CreateUser) (*int64, error) {
	userId, err := s.storage.CreateUser(userDto)
	if err != nil {
		s.logger.Error("error create user", zap.Error(err))
		return nil, err
	}

	err = s.userInfoService.CreateUserInfo(*userId, userDto)
	if err != nil {
		s.logger.Error("error create user info", zap.Error(err))
		return nil, err
	}

	err = s.userSettingsService.CreateUserSettings(*userId)
	if err != nil {
		s.logger.Error("error create user settings", zap.Error(err))
		return nil, err
	}

	return userId, nil
}

func (s *UserService) GetUserData(userId int64) (*dto.UserData, error) {
	user, err := s.storage.GetUser(userId)
	if err != nil {
		s.logger.Error("get user error", zap.Error(err))
		return nil, nil
	}

	userInfo, err := s.userInfoService.GetUserInfo(userId)
	if err != nil {
		s.logger.Error("get user info error", zap.Error(err))
		return nil, nil
	}

	userSettings, err := s.userSettingsService.GetUserSettings(userId)
	if err != nil {
		s.logger.Error("get user settings error", zap.Error(err))
		return nil, nil
	}

	return formUserDto(user, userInfo, userSettings), nil
}

func (s *UserService) GetUser(userId int64) (*models.User, error) {
	return s.storage.GetUser(userId)
}

func (s *UserService) ChangePassword(userId int64, newPasswordHash string) error {
	return s.storage.ChangePassword(userId, newPasswordHash)
}
func (s *UserService) ChangeEmail(userId int64, newEmail string) error {
	return s.storage.ChangeEmail(userId, newEmail)
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
