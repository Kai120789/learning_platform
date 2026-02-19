package service

import (
	"go.uber.org/zap"
)

type AuthService struct {
	client AuthClient
	logger *zap.Logger
}

type AuthClient interface {
}

func NewAuthService(client AuthClient, logger *zap.Logger) *AuthService {
	return &AuthService{
		client: client,
		logger: logger,
	}
}
