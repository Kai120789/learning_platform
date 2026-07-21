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
	UpdateUserInfo(userInfo dto.UserInfoRequest) error
	UpdateUserAvatar(userID int64, avatar string) error
	GetUsersShortInfo(userIDs []int64) ([]dto.UserShortInfo, error)
	UpdateUserTgUsername(userID int64, tgUsername string) error
}

func NewUserInfoService(
	storage UserInfoStorage,
) *UserInfoService {
	return &UserInfoService{
		storage: storage,
	}
}

func (ui *UserInfoService) CreateUserInfo(userID int64, userDto dto.CreateUser) error {
	err := ui.storage.CreateUserInfo(userID, userDto)
	if err != nil {
		return fmt.Errorf("create user info: %w", err)
	}

	return nil
}

func (ui *UserInfoService) GetUserInfo(userID int64) (*models.UserInfo, error) {
	userInfo, err := ui.storage.GetUserInfo(userID)
	if err != nil {
		return nil, fmt.Errorf("get user info: %w", err)
	}

	return userInfo, nil
}

func (ui *UserInfoService) UpdateUserInfo(userInfo dto.UserInfoRequest) (*models.UserInfo, error) {
	err := ui.storage.UpdateUserInfo(userInfo)
	if err != nil {
		return nil, fmt.Errorf("update user info: %w", err)
	}

	return ui.storage.GetUserInfo(userInfo.UserID)
}

func (ui *UserInfoService) UpdateUserAvatar(userID int64, avatar string) error {
	err := ui.storage.UpdateUserAvatar(userID, avatar)
	if err != nil {
		return fmt.Errorf("update user avatar: %w", err)
	}

	return nil
}

func (ui *UserInfoService) GetUsersShortInfo(userIDs []int64) ([]dto.UserShortInfo, error) {
	res, err := ui.storage.GetUsersShortInfo(userIDs)
	if err != nil {
		return nil, fmt.Errorf("get users short info: %w", err)
	}

	return res, nil
}

func (ui *UserInfoService) UpdateUserTgUsername(userID int64, tgUsername string) error {
	err := ui.storage.UpdateUserTgUsername(userID, tgUsername)
	if err != nil {
		return fmt.Errorf("update user tg link: %w", err)
	}

	return nil
}
