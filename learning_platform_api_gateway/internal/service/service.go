package service

import (
	"go.uber.org/zap"
	"learning-platform/api-gateway/internal/redis"
)

type Service struct {
	AuthService *AuthService
	UserService *UserService
}

type Client struct {
	AuthClient AuthClient
	UserClient UserClient
}

func New(client *Client, logger *zap.Logger, redis *redis.RedisStorage) *Service {
	userService := NewUserService(client.UserClient, logger)
	return &Service{
		AuthService: NewAuthService(client.AuthClient, logger, userService, redis),
		UserService: userService,
	}
}
