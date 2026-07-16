package service

import (
	"learning-platform/api-gateway/internal/dto/authDto"
	"learning-platform/api-gateway/internal/dto/enum"
	"learning-platform/api-gateway/internal/dto/userDto"
)

type UserService struct {
	client UserClient
}

type UserClient interface {
	GetUserByEmail(email string) (*userDto.GetUser, error)
	GetUserById(id int64) (*userDto.GetUser, error)
	GetUserData(id int64) (*userDto.UserData, error)
	CreateUser(newUser authDto.RegisterRequest) (*int64, error)
	UpdateUserInfo(userID int64, userInfo userDto.UserInfoRequest) (*userDto.UserInfoResponse, error)
	UpdateUserSettings(userID int64, userSettings userDto.UserSettingsRequest) (*userDto.UserSettingsResponse, error)
	UpdateUserTheme(userID int64, theme enum.UserTheme) error
	UpdateUserAvatar(userID int64, avatar string) error
}

func NewUserService(client UserClient) *UserService {
	return &UserService{
		client: client,
	}
}

func (u *UserService) GetUserByEmail(email string) (*userDto.GetUser, error) {
	res, err := u.client.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (u *UserService) GetUserById(id int64) (*userDto.GetUser, error) {
	res, err := u.client.GetUserById(id)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (u *UserService) GetUserData(id int64) (*userDto.UserData, error) {
	res, err := u.client.GetUserData(id)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (u *UserService) CreateUser(newUser authDto.RegisterRequest) (*int64, error) {
	res, err := u.client.CreateUser(newUser)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (u *UserService) UpdateUserInfo(userID int64, userInfo userDto.UserInfoRequest) (*userDto.UserInfoResponse, error) {
	res, err := u.client.UpdateUserInfo(userID, userInfo)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (u *UserService) UpdateUserSettings(userID int64, userSettings userDto.UserSettingsRequest) (*userDto.UserSettingsResponse, error) {
	res, err := u.client.UpdateUserSettings(userID, userSettings)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (u *UserService) UpdateUserTheme(userID int64, theme enum.UserTheme) error {
	err := u.client.UpdateUserTheme(userID, theme)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserService) UpdateUserAvatar(userID int64, avatar string) error {
	err := u.client.UpdateUserAvatar(userID, avatar)
	if err != nil {
		return err
	}

	return nil
}
