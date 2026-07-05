package service

type LessonService struct {
	client LessonClient
}

type LessonClient interface{}
