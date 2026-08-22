package service

import "learning-platform/tutors/internal/dto"

type TutorSubjectService struct {
	storage TutorSubjectStorage
}

type TutorSubjectStorage interface {
	AddTutorSubjects(tutorID int64, subjectIDs []int64) ([]int64, error)
	UpdateTutorSubjects(tutorID int64, subjectIDs []int64) ([]int64, error)
	GetOneTutorSubjectIDs(tutorID int64) ([]int64, error)
	GetTutorsSubjectIDs(tutorIDs []int64) ([][]int64, error)
	GetTutorsBySubject(getTutors dto.GetTutors) ([]int64, int64, error)
}

func NewTutorSubjectService(storage TutorSubjectStorage) *TutorSubjectService {
	return &TutorSubjectService{
		storage: storage,
	}
}

func (tsu *TutorSubjectService) AddTutorSubjects(tutorID int64, subjectIDs []int64) ([]int64, error) {
	res, err := tsu.storage.AddTutorSubjects(tutorID, subjectIDs)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (tsu *TutorSubjectService) UpdateTutorSubjects(tutorID int64, subjectIDs []int64) ([]int64, error) {
	res, err := tsu.storage.UpdateTutorSubjects(tutorID, subjectIDs)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (tsu *TutorSubjectService) GetOneTutorSubjectIDs(tutorID int64) ([]int64, error) {
	res, err := tsu.storage.GetOneTutorSubjectIDs(tutorID)
	if err != nil {
		return nil, err
	}

	return res, err
}

func (tsu *TutorSubjectService) GetTutorsSubjectIDs(tutorIDs []int64) ([][]int64, error) {
	res, err := tsu.storage.GetTutorsSubjectIDs(tutorIDs)
	if err != nil {
		return nil, err
	}

	return res, err
}

func (tsu *TutorSubjectService) GetTutorsBySubject(getTutors dto.GetTutors) ([]int64, int64, error) {
	res, count, err := tsu.storage.GetTutorsBySubject(getTutors)
	if err != nil {
		return nil, 0, err
	}

	return res, count, err
}
