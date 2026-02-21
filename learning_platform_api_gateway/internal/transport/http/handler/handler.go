package handler

import "go.uber.org/zap"

type Handler struct {
	AuthHandler AuthHandler
	UserHandler UserHandler
}

type Service struct {
	AuthService AuthService
	UserService UserService
}

func New(service *Service, logger *zap.Logger) *Handler {
	return &Handler{
		AuthHandler: *NewAuthHandler(service.AuthService, logger),
		UserHandler: *NewUserHandler(service.UserService, logger),
	}
}
