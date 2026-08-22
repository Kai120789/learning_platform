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
	MaterialHandler *MaterialHandler
	TutorHandler    *TutorHandler
}

type Service struct {
	AuthService     AuthService
	UserService     UserService
	GroupService    GroupService
	LessonService   LessonService
	SubjectService  SubjectService
	ScheduleService ScheduleService
	MaterialService MaterialService
	TutorService    TutorService
}

func New(service *Service, logger *zap.Logger, cfg *config.Config) *Handler {
	return &Handler{
		AuthHandler:     NewAuthHandler(service.AuthService, logger, cfg),
		UserHandler:     NewUserHandler(service.UserService, logger),
		GroupHandler:    NewGroupHandler(service.GroupService, logger),
		LessonHandler:   NewLessonHandler(service.LessonService, logger),
		SubjectHandler:  NewSubjectHandler(service.SubjectService, logger),
		ScheduleHandler: NewScheduleHandler(service.ScheduleService, logger),
		MaterialHandler: NewMaterialHandler(service.MaterialService, logger),
		TutorHandler:    NewTutorHandler(service.TutorService, logger),
	}
}
