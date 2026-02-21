package handler

import "go.uber.org/zap"

type UserHandler struct {
	service UserService
	logger  *zap.Logger
}

type UserService interface {
}

func NewUserHandler(service UserService, logger *zap.Logger) *UserHandler {
	return &UserHandler{
		service: service,
		logger:  logger,
	}
}
