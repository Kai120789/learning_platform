package service

import "learning-platform/api-gateway/internal/dto/subjectDto"

type SubjectService struct {
	client SubjectClient
}

type SubjectClient interface {
	GetOneSubject(subjectID int64) (*subjectDto.Subject, error)
	GetAllSubjects() ([]subjectDto.Subject, error)
	GetUserSubjects(userID int64) ([]subjectDto.Subject, error)
	SetUserSubjects(userID int64, subjectIDs []int64) error
	UpdateUserSubjects(userID int64, subjectIDs, deletedSubjectIDs []int64) error
}

func NewSubjectService(client SubjectClient) *SubjectService {
	return &SubjectService{
		client: client,
	}
}

func (s *SubjectService) GetOneSubject(subjectID int64) (*subjectDto.Subject, error) {
	res, err := s.client.GetOneSubject(subjectID)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *SubjectService) GetAllSubjects() ([]subjectDto.Subject, error) {
	res, err := s.client.GetAllSubjects()
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *SubjectService) GetUserSubjects(userID int64) ([]subjectDto.Subject, error) {
	res, err := s.client.GetUserSubjects(userID)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *SubjectService) SetUserSubjects(userID int64, subjectIDs []int64) error {
	err := s.client.SetUserSubjects(userID, subjectIDs)
	if err != nil {
		return err
	}

	return nil
}

func (s *SubjectService) UpdateUserSubjects(userID int64, subjectIDs, deletedSubjectIDs []int64) error {
	err := s.client.UpdateUserSubjects(userID, subjectIDs, deletedSubjectIDs)
	if err != nil {
		return err
	}

	return nil
}
