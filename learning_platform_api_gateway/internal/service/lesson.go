package service

import (
	"learning-platform/api-gateway/internal/dto/enum"
	"learning-platform/api-gateway/internal/dto/lessonDto"
)

type LessonService struct {
	client LessonClient
}

type LessonClient interface {
	GetOneLesson(lessonID int64) (*lessonDto.LessonResponse, error)
	GetLessonsByUserId(userID int64) ([]lessonDto.LessonResponse, error)
	CreateLesson(lesson lessonDto.CreateLesson) (*lessonDto.LessonResponse, error)
	UpdateLesson(lesson lessonDto.UpdateLesson) (*lessonDto.LessonResponse, error)
	UpdateLessonStatus(lessonID int64, status enum.LessonStatus) error
	GetLessonsByTutorId(tutorID int64) ([]lessonDto.LessonResponse, error)
}

func NewLessonService(client LessonClient) *LessonService {
	return &LessonService{
		client: client,
	}
}

func (l *LessonService) GetOneLesson(lessonID int64) (*lessonDto.LessonResponse, error) {
	res, err := l.client.GetOneLesson(lessonID)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (l *LessonService) GetLessonsByUserId(userID int64) ([]lessonDto.LessonResponse, error) {
	res, err := l.client.GetLessonsByUserId(userID)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (l *LessonService) CreateLesson(lesson lessonDto.CreateLesson) (*lessonDto.LessonResponse, error) {
	res, err := l.client.CreateLesson(lesson)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (l *LessonService) UpdateLesson(lesson lessonDto.UpdateLesson) (*lessonDto.LessonResponse, error) {
	res, err := l.client.UpdateLesson(lesson)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (l *LessonService) UpdateLessonStatus(lessonID int64, status enum.LessonStatus) error {
	err := l.client.UpdateLessonStatus(lessonID, status)
	if err != nil {
		return err
	}

	return nil
}

func (l *LessonService) GetLessonsByTutorId(tutorID int64) ([]lessonDto.LessonResponse, error) {
	res, err := l.client.GetLessonsByTutorId(tutorID)
	if err != nil {
		return nil, err
	}

	return res, nil
}
