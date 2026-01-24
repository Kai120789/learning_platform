package service

import (
	"go.uber.org/zap"
	"learning-platform/auth/internal/config"
	"learning-platform/auth/internal/storage/postgres"
	"learning-platform/auth/internal/storage/redis"
)

type AuthService struct {
	config   *config.Config
	logger   *zap.Logger
	postgres *postgres.PostgresStorage
	redis    *redis.RedisStorage
}

type AuthStorage interface {
}

func New(
	config *config.Config,
	logger *zap.Logger,
	postgres *postgres.PostgresStorage,
	redis *redis.RedisStorage,
) *AuthService {
	return &AuthService{
		config:   config,
		logger:   logger,
		postgres: postgres,
		redis:    redis,
	}
}

func (*AuthService) Login() {}

func (*AuthService) Register() {}

func (*AuthService) RefreshTokens() {}

func (*AuthService) Logout() {}

func (*AuthService) LogoutAll() {}

func (*AuthService) ChangePassword() {}

func (*AuthService) ForceChangePassword() {}

func (*AuthService) ChangeEmail() {}

func (*AuthService) ForceChangeEmail() {}
