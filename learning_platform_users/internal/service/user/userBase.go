package user

import (
	"github.com/Kai120789/learning_platform_models/models"
	"go.uber.org/zap"
	"learning-platform/users/internal/dto"
)

type UserBaseService struct {
	logger              *zap.Logger
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
	logger *zap.Logger,
	storage UserBaseStorage,
	userInfoService UserInfo,
	userSettingsService UserSettings,
) *UserBaseService {
	return &UserBaseService{
		logger:              logger,
		storage:             storage,
		userInfoService:     userInfoService,
		userSettingsService: userSettingsService,
	}
}

func (s *UserBaseService) CreateUser(userDto dto.CreateUser) (*int64, error) {
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

func (s *UserBaseService) GetUserData(userId int64) (*dto.UserData, error) {
	user, err := s.storage.GetUserById(userId)
	if err != nil {
		s.logger.Error("get user error", zap.Error(err))
		return nil, err
	}

	userInfo, err := s.userInfoService.GetUserInfo(userId)
	if err != nil {
		s.logger.Error("get user info error", zap.Error(err))
		return nil, err
	}

	userSettings, err := s.userSettingsService.GetUserSettings(userId)
	if err != nil {
		s.logger.Error("get user settings error", zap.Error(err))
		return nil, err
	}

	return formUserDto(user, userInfo, userSettings), nil
}

func (s *UserBaseService) GetUserById(userId int64) (*models.User, error) {
	return s.storage.GetUserById(userId)
}

func (s *UserBaseService) GetUserByEmail(email string) (*models.User, error) {
	return s.storage.GetUserByEmail(email)
}

func (s *UserBaseService) ChangePassword(userId int64, newPasswordHash string) error {
	return s.storage.ChangePassword(userId, newPasswordHash)
}
func (s *UserBaseService) ChangeEmail(userId int64, newEmail string) error {
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
