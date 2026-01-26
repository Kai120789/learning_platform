package service

import "go.uber.org/zap"

type UserService struct {
	logger  *zap.Logger
	storage *UserStorage
}

type UserStorage interface {
}

func NewUserService(
	logger *zap.Logger,
	storage *UserStorage,
) *UserService {
	return &UserService{
		logger:  logger,
		storage: storage,
	}
}
