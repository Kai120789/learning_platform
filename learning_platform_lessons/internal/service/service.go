package service

import (
	"fmt"
	"learning-platform/lessons/internal/dto"
	"learning-platform/lessons/internal/models"
	"learning-platform/lessons/internal/models/enum"
	"learning-platform/lessons/internal/utils"
)

type Service struct {
	storage *Storage
}

type Storage struct {
	LessonStorage      LessonStorage
	LessonMediaStorage LessonMediaStorage
	LessonUserStorage  LessonUserStorage
}

type LessonStorage interface {
	CreateLesson(lessonDto dto.CreateLesson) (*models.Lesson, error)
	GetAllLesson() ([]models.Lesson, error)
	GetLessonById(lessonID int64) (*models.Lesson, error)
	GetLessonsByUserId(userID int64) ([]models.Lesson, error)
	GetLessonsByTutorId(tutorID int64) ([]models.Lesson, error)
	UpdateLesson(lessonDto dto.UpdateLesson) (*models.Lesson, error)
	UpdateLessonStatus(lessonID int64, status enum.LessonStatus) error
}

type LessonMediaStorage interface {
	GetAllLessonMedias(lessonID int64) ([]models.LessonMedia, error)
	SetLessonMedias(lessonID int64, mediaIDs []dto.MediaItem) error
	DeleteLessonMedias(lessonID int64, mediaIDs []int64) error
}

type LessonUserStorage interface {
	SetUsersToLesson(lessonID int64, userIDs []int64) error
	DeleteUsersFromLesson(lessonID int64, userIDs []int64) error
	GetAllLessonParticipants(lessonID int64) ([]int64, error)
}

func New(storage *Storage) *Service {
	return &Service{
		storage: storage,
	}
}

func (l *Service) GetOneLesson(lessonID int64) (*dto.LessonResponse, error) {
	resLesson, err := l.storage.LessonStorage.GetLessonById(lessonID)
	if err != nil {
		return nil, fmt.Errorf("get one lesson %d: %w", lessonID, err)
	}

	lessonWithMedias, err := l.buildOneLessonWithMediasDto(resLesson)
	if err != nil {
		return nil, fmt.Errorf("get one lesson %d (get lesson medias): %w", lessonID, err)
	}

	return lessonWithMedias, nil
}

func (l *Service) GetLessonsByUserId(userID int64) ([]dto.LessonResponse, error) {
	var resLessons []dto.LessonResponse
	lessons, err := l.storage.LessonStorage.GetLessonsByUserId(userID)
	if err != nil {
		return nil, fmt.Errorf("get lessons by user id %d: %w", userID, err)
	}

	for _, oneLesson := range lessons {
		oneLessonWithMedias, err := l.buildOneLessonWithMediasDto(&oneLesson)
		if err != nil {
			return nil, fmt.Errorf("get lessons by user id %d (get lesson medias): %w", userID, err)
		}

		resLessons = append(resLessons, *oneLessonWithMedias)
	}

	return resLessons, nil
}

func (l *Service) CreateLesson(lesson dto.CreateLesson) (*dto.LessonResponse, error) {
	resLesson, err := l.storage.LessonStorage.CreateLesson(lesson)
	if err != nil {
		return nil, fmt.Errorf("create lesson: %w", err)
	}

	err = l.storage.LessonMediaStorage.SetLessonMedias(resLesson.ID, lesson.MediaItems)
	if err != nil {
		return nil, fmt.Errorf("create lesson (set lesson medias): %w", err)
	}

	err = l.storage.LessonUserStorage.SetUsersToLesson(resLesson.ID, lesson.UserIDs)
	if err != nil {
		return nil, fmt.Errorf("create lesson (set users to lesson): %w", err)
	}

	lessonWithMedias, err := l.buildOneLessonWithMediasDto(resLesson)
	if err != nil {
		return nil, fmt.Errorf("create lesson (get lesson medias): %w", err)
	}

	return lessonWithMedias, nil
}

func (l *Service) UpdateLesson(lesson dto.UpdateLesson) (*dto.LessonResponse, error) {
	resLesson, err := l.storage.LessonStorage.UpdateLesson(lesson)
	if err != nil {
		return nil, fmt.Errorf("update lesson %d: %w", lesson.ID, err)
	}

	err = l.storage.LessonMediaStorage.DeleteLessonMedias(lesson.ID, lesson.DeletedMediaIDs)
	if err != nil {
		return nil, fmt.Errorf("update lesson %d (delete lesson medias): %w", lesson.ID, err)
	}

	err = l.storage.LessonMediaStorage.SetLessonMedias(resLesson.ID, lesson.MediaItems)
	if err != nil {
		return nil, fmt.Errorf("update lesson %d (set lesson medias): %w", lesson.ID, err)
	}

	err = l.storage.LessonUserStorage.SetUsersToLesson(resLesson.ID, lesson.UserIDs)
	if err != nil {
		return nil, fmt.Errorf("update lesson %d (set users to lesson): %w", lesson.ID, err)
	}

	err = l.storage.LessonUserStorage.DeleteUsersFromLesson(resLesson.ID, lesson.DeletedUserIDs)
	if err != nil {
		return nil, fmt.Errorf("update lesson %d (delete users from lesson): %w", lesson.ID, err)
	}

	lessonWithMedias, err := l.buildOneLessonWithMediasDto(resLesson)
	if err != nil {
		return nil, fmt.Errorf("update lesson %d (get lesson medias): %w", lesson.ID, err)
	}

	return lessonWithMedias, nil
}

func (l *Service) UpdateLessonStatus(lessonID int64, status enum.LessonStatus) error {
	err := l.storage.LessonStorage.UpdateLessonStatus(lessonID, status)
	if err != nil {
		return fmt.Errorf("update lesson %d status to %s: %w", lessonID, status, err)
	}

	return nil
}

func (l *Service) GetLessonsByTutorId(tutorID int64) ([]dto.LessonResponse, error) {
	var resLessons []dto.LessonResponse
	lessons, err := l.storage.LessonStorage.GetLessonsByTutorId(tutorID)
	if err != nil {
		return nil, fmt.Errorf("get lessons by tutor id %d: %w", tutorID, err)
	}

	for _, oneLesson := range lessons {
		oneLessonWithMedias, err := l.buildOneLessonWithMediasDto(&oneLesson)
		if err != nil {
			return nil, fmt.Errorf("get lessons by tutor id %d (get lesson medias): %w", tutorID, err)
		}

		resLessons = append(resLessons, *oneLessonWithMedias)
	}

	return resLessons, nil
}

func (l *Service) buildOneLessonWithMediasDto(
	lesson *models.Lesson,
) (*dto.LessonResponse, error) {
	lessonMedias, err := l.storage.LessonMediaStorage.GetAllLessonMedias(lesson.ID)
	if err != nil {
		return nil, fmt.Errorf("get lesson %d medias: %w", lesson.ID, err)
	}

	lessonUsers, err := l.storage.LessonUserStorage.GetAllLessonParticipants(lesson.ID)
	if err != nil {
		return nil, fmt.Errorf("get lesson %d participants: %w", lesson.ID, err)
	}

	var resLessonMedias []dto.MediaItemResponse
	for _, oneMedia := range lessonMedias {
		resLessonMedias = append(
			resLessonMedias,
			dto.MediaItemResponse{
				ID:        oneMedia.ID,
				LessonID:  oneMedia.LessonID,
				S3Link:    oneMedia.S3Link,
				S3Preview: oneMedia.S3Preview,
				Type:      oneMedia.Type,
			},
		)
	}

	return &dto.LessonResponse{
		ID:         lesson.ID,
		BoardID:    utils.DBInt8ToOptional(lesson.BoardID),
		MeetLink:   utils.DBStringToOptional(lesson.MeetLink),
		StartTime:  lesson.StartTime.Time,
		Duration:   lesson.Duration,
		TutorID:    lesson.TutorID,
		Status:     lesson.Status,
		UserIDs:    lessonUsers,
		MediaItems: resLessonMedias,
	}, nil
}
