package service

import (
	"fmt"
	"learning-platform/users/internal/dto"
	"learning-platform/users/internal/models"
)

type UserInfoService struct {
	storage UserInfoStorage
}

type UserInfoStorage interface {
	CreateUserInfo(userID int64, userDto dto.CreateUser) error
	GetUserInfo(userID int64) (*models.UserInfo, error)
	UpdateUserInfo(userInfo dto.UserInfo) error
}

func NewUserInfoService(
	storage UserInfoStorage,
) *UserInfoService {
	return &UserInfoService{
		storage: storage,
	}
}

func (s *UserInfoService) CreateUserInfo(userID int64, userDto dto.CreateUser) error {
	err := s.storage.CreateUserInfo(userID, userDto)
	if err != nil {
		return fmt.Errorf("create user info: %w", err)
	}

	return nil
}

func (s *UserInfoService) GetUserInfo(userID int64) (*models.UserInfo, error) {
	userInfo, err := s.storage.GetUserInfo(userID)
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

	return s.storage.GetUserInfo(userInfo.UserID)
}
