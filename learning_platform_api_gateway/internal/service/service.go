package service

import (
	"learning-platform/api-gateway/internal/redis"
)

type Service struct {
	AuthService     *AuthService
	UserService     *UserService
	GroupService    *GroupService
	LessonService   *LessonService
	ScheduleService *ScheduleService
	SubjectService  *SubjectService
}

type Client struct {
	AuthClient     AuthClient
	UserClient     UserClient
	GroupClient    GroupClient
	LessonClient   LessonClient
	ScheduleClient ScheduleClient
	SubjectClient  SubjectClient
}

func New(client *Client, redis *redis.RedisStorage) *Service {
	userService := NewUserService(client.UserClient)
	return &Service{
		AuthService:     NewAuthService(client.AuthClient, userService, redis),
		UserService:     userService,
		GroupService:    NewGroupService(client.GroupClient),
		LessonService:   NewLessonService(client.LessonClient),
		ScheduleService: NewScheduleService(client.ScheduleClient),
		SubjectService:  NewSubjectService(client.SubjectClient),
	}
}
