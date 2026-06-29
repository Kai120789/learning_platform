package service

import "learning-platform/subjects/internal/models"

type SubjectService struct {
	storage SubjectStorage
}

type SubjectStorage interface {
	GetOneSubject(subjectID int64) (*models.Subject, error)
	GetAllSubjects() ([]models.Subject, error)
}

func NewSubjectService(storage SubjectStorage) *SubjectService {
	return &SubjectService{
		storage: storage,
	}
}

func (s *SubjectService) GetOneSubject(subjectID int64) (*models.Subject, error) {
	res, err := s.storage.GetOneSubject(subjectID)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *SubjectService) GetAllSubjects() ([]models.Subject, error) {
	res, err := s.storage.GetAllSubjects()
	if err != nil {
		return nil, err
	}
	return res, nil
}
