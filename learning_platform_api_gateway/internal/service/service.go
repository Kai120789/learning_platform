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
	MediaService    *MediaService
	MaterialService *MaterialService
	TutorService    *TutorService
}

type Client struct {
	AuthClient     AuthClient
	UserClient     UserClient
	GroupClient    GroupClient
	LessonClient   LessonClient
	ScheduleClient ScheduleClient
	SubjectClient  SubjectClient
	MediaClient    MediaClient
	MaterialClient MaterialClient
	TutorClient    TutorClient
}

func New(client *Client, redis *redis.RedisStorage) *Service {
	userService := NewUserService(client.UserClient)
	subjectService := NewSubjectService(client.SubjectClient)
	mediaService := NewMediaService(client.MediaClient)
	return &Service{
		AuthService:     NewAuthService(client.AuthClient, userService, redis),
		UserService:     userService,
		GroupService:    NewGroupService(client.GroupClient, userService, subjectService),
		LessonService:   NewLessonService(client.LessonClient),
		ScheduleService: NewScheduleService(client.ScheduleClient),
		SubjectService:  subjectService,
		MediaService:    mediaService,
		MaterialService: NewMaterialService(client.MaterialClient, mediaService),
		TutorService:    NewTutorService(client.TutorClient, subjectService, userService),
	}
}
