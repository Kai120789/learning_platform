package service

import (
	"go.uber.org/zap"
	"learning-platform/users/internal/dto"
)

type RegisterUseCase struct {
	logger  *zap.Logger
	storage RegisterStorage
}

type RegisterStorage interface {
	CreateUser(dto dto.CreateUser) (*int64, error)
	CreateUserInfo(userID int64, dto dto.CreateUser) error
	CreateUserSettings(userID int64) error
}

func NewRegisterUseCase(
	logger *zap.Logger,
	storage RegisterStorage,
) *RegisterUseCase {
	return &RegisterUseCase{
		logger:  logger,
		storage: storage,
	}
}

func (s *RegisterUseCase) CreateUser(userDto dto.CreateUser) (*int64, error) {
	userId, err := s.storage.CreateUser(userDto)
	if err != nil {
		s.logger.Error("error create user", zap.Error(err))
		return nil, err
	}

	err = s.storage.CreateUserInfo(*userId, userDto)
	if err != nil {
		s.logger.Error("error create user info", zap.Error(err))
		return nil, err
	}

	err = s.storage.CreateUserSettings(*userId)
	if err != nil {
		s.logger.Error("error create user settings", zap.Error(err))
		return nil, err
	}

	return userId, nil
}
