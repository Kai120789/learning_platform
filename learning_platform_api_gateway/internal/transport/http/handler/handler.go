package handler

import (
	"go.uber.org/zap"
	"learning-platform/api-gateway/internal/config"
)

type Handler struct {
	AuthHandler *AuthHandler
	UserHandler *UserHandler
}

type Service struct {
	AuthService AuthService
	UserService UserService
}

func New(service *Service, logger *zap.Logger, cfg *config.Config) *Handler {
	return &Handler{
		AuthHandler: NewAuthHandler(service.AuthService, logger, cfg),
		UserHandler: NewUserHandler(service.UserService, logger),
	}
}
