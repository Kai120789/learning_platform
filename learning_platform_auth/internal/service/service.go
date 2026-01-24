package service

import (
	"go.uber.org/zap"
	"learning-platform/auth/internal/config"
)

type AuthService struct {
	config  *config.Config
	logger  *zap.Logger
	storage *AuthStorage
}

type AuthStorage interface {
}

func New(config *config.Config, logger *zap.Logger, storage *AuthStorage) *AuthService {
	return &AuthService{
		config:  config,
		logger:  logger,
		storage: storage,
	}
}

func (*AuthService) Login() {}

func (*AuthService) Register() {}

func (*AuthService) RefreshTokens() {}

func (*AuthService) Logout() {}

func (*AuthService) LogoutAll() {}

func (*AuthService) ChangePassword() {}

func (*AuthService) ForceChangePassword() {}
