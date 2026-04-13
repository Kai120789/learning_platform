package user

import (
	"fmt"
	"github.com/Kai120789/learning_platform_models/models"
	"learning-platform/users/internal/dto"
)

type UserInfoService struct {
	storage UserInfoStorage
}

type UserInfoStorage interface {
	CreateUserInfo(userId int64, userDto dto.CreateUser) error
	GetUserInfo(userId int64) (*models.UserInfo, error)
	UpdateUserInfo(userInfo dto.UserInfo) error
}

func NewUserInfoService(
	storage UserInfoStorage,
) *UserInfoService {
	return &UserInfoService{
		storage: storage,
	}
}

func (s *UserInfoService) CreateUserInfo(userId int64, userDto dto.CreateUser) error {
	err := s.storage.CreateUserInfo(userId, userDto)
	if err != nil {
		return fmt.Errorf("create user info: %w", err)
	}

	return nil
}

func (s *UserInfoService) GetUserInfo(userId int64) (*models.UserInfo, error) {
	userInfo, err := s.storage.GetUserInfo(userId)
	if err != nil {
		return nil, fmt.Errorf("get user info: %w", err)
	}

	return userInfo, nil
}

func (s *UserInfoService) UpdateUserInfo(userInfo dto.UserInfo) (*models.UserInfo, error) {
	err := s.storage.UpdateUserInfo(userInfo)
	if err != nil {
		return nil, fmt.Errorf("update user info: %w", err)
	}

	return s.storage.GetUserInfo(userInfo.UserId)
}
