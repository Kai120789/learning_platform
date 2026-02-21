package handler

import "go.uber.org/zap"

type AuthHandler struct {
	service AuthService
	logger  *zap.Logger
}

type AuthService interface {
}

func NewAuthHandler(service AuthService, logger *zap.Logger) *AuthHandler {
	return &AuthHandler{
		service: service,
		logger:  logger,
	}
}
