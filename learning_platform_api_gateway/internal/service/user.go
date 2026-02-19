package service

import (
	"go.uber.org/zap"
)

type UserService struct {
	client UserClient
	logger *zap.Logger
}

type UserClient interface {
}

func NewUserService(client UserClient, logger *zap.Logger) *UserService {
	return &UserService{
		client: client,
		logger: logger,
	}
}
