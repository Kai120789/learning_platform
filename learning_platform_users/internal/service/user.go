package service

import (
	"go.uber.org/zap"
	"learning-platform/users/internal/dto"
)

type UserService struct {
	logger              *zap.Logger
	storage             UserStorage
	userInfoService     UserInfoService
	userSettingsService UserSettingsService
}

type UserStorage interface {
	CreateUser(userDto dto.CreateUser) (*int64, error)
	GetUser(userId int64) (*dto.GetUser, error)
}

func NewUserService(
	logger *zap.Logger,
	storage UserStorage,
	userInfoService UserInfoService,
	userSettingsService UserSettingsService,
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

	err = s.userInfoService.storage.CreateUserInfo(*userId, userDto)
	if err != nil {
		s.logger.Error("error create user info", zap.Error(err))
		return nil, err
	}

	err = s.userSettingsService.storage.CreateUserSettings(*userId)
	if err != nil {
		s.logger.Error("error create user settings", zap.Error(err))
		return nil, err
	}

	return userId, nil
}

func (s *UserService) GetUser(userId int64) (*dto.GetUser, error) {
	return s.storage.GetUser(userId)
}
