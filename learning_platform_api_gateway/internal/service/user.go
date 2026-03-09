package service

import (
	"go.uber.org/zap"
	"learning-platform/api-gateway/internal/dto"
)

type UserService struct {
	client UserClient
	logger *zap.Logger
}

type UserClient interface {
	GetUserByEmail(email string) (*dto.GetUser, error)
	GetUserById(id int64) (*dto.GetUser, error)
	GetUserData(id int64) (*dto.UserData, error)
	CreateUser(newUser dto.RegisterRequest) (*int64, error)
}

func NewUserService(client UserClient, logger *zap.Logger) *UserService {
	return &UserService{
		client: client,
		logger: logger,
	}
}

func (u *UserService) GetUserByEmail(email string) (*dto.GetUser, error) {
	res, err := u.client.GetUserByEmail(email)
	if err != nil {
		u.logger.Error("failed get user by email", zap.Error(err))
		return nil, err
	}

	return res, nil
}

func (u *UserService) GetUserById(id int64) (*dto.GetUser, error) {
	res, err := u.client.GetUserById(id)
	if err != nil {
		u.logger.Error("failed get user by id", zap.Error(err))
		return nil, err
	}

	return res, nil
}

func (u *UserService) GetUserData(id int64) (*dto.UserData, error) {
	res, err := u.client.GetUserData(id)
	if err != nil {
		u.logger.Error("failed get user data", zap.Error(err))
		return nil, err
	}

	return res, nil
}

func (u *UserService) CreateUser(newUser dto.RegisterRequest) (*int64, error) {
	res, err := u.client.CreateUser(newUser)
	if err != nil {
		u.logger.Error("failed create user", zap.Error(err))
		return nil, err
	}

	return res, nil
}
