package service

import (
	"go.uber.org/zap"
	"learning-platform/users/internal/dto"
)

type UserService struct {
	logger  *zap.Logger
	storage UserStorage
}

type UserStorage interface {
	CreateUser(userDto dto.CreateUser) (*int64, error)
}

func NewUserService(
	logger *zap.Logger,
	storage UserStorage,
) *UserService {
	return &UserService{
		logger:  logger,
		storage: storage,
	}
}
