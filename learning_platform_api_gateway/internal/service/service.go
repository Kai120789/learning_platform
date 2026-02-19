package service

import "go.uber.org/zap"

type Service struct {
	AuthService *AuthService
	UserService *UserService
}

type Client struct {
	AuthClient AuthClient
	UserClient UserClient
}

func New(client *Client, logger *zap.Logger) *Service {
	return &Service{
		AuthService: NewAuthService(client.AuthClient, logger),
		UserService: NewUserService(client.UserClient, logger),
	}
}
