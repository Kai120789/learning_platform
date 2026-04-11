package service

import (
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

func New(client *Client, redis *redis.RedisStorage) *Service {
	userService := NewUserService(client.UserClient)
	return &Service{
		AuthService: NewAuthService(client.AuthClient, userService, redis),
		UserService: userService,
	}
}
