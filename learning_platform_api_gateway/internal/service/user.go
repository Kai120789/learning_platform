package service

import (
	"learning-platform/api-gateway/internal/dto/authDto"
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
