package service

import "go.uber.org/zap"

type UserInfoService struct {
	logger  *zap.Logger
	storage *UserInfoStorage
}

type UserInfoStorage interface {
}

func NewUserInfoService(
	logger *zap.Logger,
	storage *UserInfoStorage,
) *UserInfoService {
	return &UserInfoService{
		logger:  logger,
		storage: storage,
	}
}
