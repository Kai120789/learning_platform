package service

import (
	"github.com/Kai120789/learning_platform_models/models"
	"go.uber.org/zap"
	"learning-platform/users/internal/dto"
)

type UserInfoService struct {
	logger  *zap.Logger
	storage UserInfoStorage
}

type UserInfoStorage interface {
	CreateUserInfo(userId int64, userDto dto.CreateUser) error
	GetUserInfo(userId int64) (*models.UserInfo, error)
	UpdateUserInfo(userInfo dto.UserInfo) error
}

func NewUserInfoService(
	logger *zap.Logger,
	storage UserInfoStorage,
) *UserInfoService {
	return &UserInfoService{
		logger:  logger,
		storage: storage,
	}
}

func (s *UserInfoService) CreateUserInfo(userId int64, userDto dto.CreateUser) error {
	return s.storage.CreateUserInfo(userId, userDto)
}

func (s *UserInfoService) GetUserInfo(userId int64) (*models.UserInfo, error) {
	return s.storage.GetUserInfo(userId)
}

func (s *UserInfoService) UpdateUserInfo(userInfo dto.UserInfo) (*models.UserInfo, error) {
	err := s.storage.UpdateUserInfo(userInfo)
	if err != nil {
		s.logger.Error("update user info error", zap.Error(err))
		return nil, err
	}

	return s.storage.GetUserInfo(userInfo.UserId)
}
