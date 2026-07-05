package handler

import (
	"go.uber.org/zap"
	"learning-platform/api-gateway/internal/config"
)

type Handler struct {
	AuthHandler     *AuthHandler
	UserHandler     *UserHandler
	GroupHandler    *GroupHandler
	LessonHandler   *LessonHandler
	SubjectHandler  *SubjectHandler
	ScheduleHandler *ScheduleHandler
}

type Service struct {
	AuthService     AuthService
	UserService     UserService
	GroupService    GroupService
	LessonService   LessonService
	SubjectService  SubjectService
	ScheduleService ScheduleService
}

func New(service *Service, logger *zap.Logger, cfg *config.Config) *Handler {
	return &Handler{
		AuthHandler:     NewAuthHandler(service.AuthService, logger, cfg),
		UserHandler:     NewUserHandler(service.UserService, logger),
		GroupHandler:    NewGroupHandler(service.GroupService, logger),
		LessonHandler:   NewLessonHandler(service.LessonService, logger),
		SubjectHandler:  NewSubjectHandler(service.SubjectService, logger),
		ScheduleHandler: NewScheduleHandler(service.ScheduleService, logger),
	}
}
