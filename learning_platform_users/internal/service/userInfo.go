package service

import (
	"go.uber.org/zap"
	"learning-platform/users/internal/dto"
)

type UserInfoService struct {
	logger  *zap.Logger
	storage UserInfoStorage
}

type UserInfoStorage interface {
	CreateUserInfo(userId int64, userDto dto.CreateUser) error
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
