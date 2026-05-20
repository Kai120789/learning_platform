package service

import (
	"learning-platform/api-gateway/internal/redis"
)

type Service struct {
	AuthService  *AuthService
	UserService  *UserService
	GroupService *GroupService
}

type Client struct {
	AuthClient  AuthClient
	UserClient  UserClient
	GroupClient GroupClient
}

func New(client *Client, redis *redis.RedisStorage) *Service {
	userService := NewUserService(client.UserClient)
	return &Service{
		AuthService:  NewAuthService(client.AuthClient, userService, redis),
		UserService:  userService,
		GroupService: NewGroupService(client.GroupClient),
	}
}
