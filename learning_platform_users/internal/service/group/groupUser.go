package group

import "go.uber.org/zap"

type GroupUserService struct {
	logger  *zap.Logger
	storage GroupUserStorage
}

type GroupUserStorage interface {
}

func NewGroupUserService(
	logger *zap.Logger,
	storage GroupUserStorage,
) *GroupUserService {
	return &GroupUserService{
		logger:  logger,
		storage: storage,
	}
}
