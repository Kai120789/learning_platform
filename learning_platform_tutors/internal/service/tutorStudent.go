package service

import (
	"learning-platform/tutors/internal/dto"
	"learning-platform/tutors/internal/models"
)

type TutorStudentService struct {
	storage TutorStudentStorage
}

type TutorStudentStorage interface {
	AddStudent(tutorID, studentID int64) error
	DeleteOneStudent(tutorID, studentID int64) error
	DeleteStudents(tutorID int64, studentIDs []int64) error
	GetTutorStudents(getTutorStudents dto.GetTutorStudents) ([]models.TutorStudent, int64, error)
	UpdateLastInteracted(tutorID, studentID int64) error
	GetOneTutorStudentsCount(tutorID int64, interactedWithinDays *int64) (int64, error)
	GetTutorsStudentsCount(tutorIDs []int64) ([]int64, error)
}

func NewTutorStudentService(storage TutorStudentStorage) *TutorStudentService {
	return &TutorStudentService{
		storage: storage,
	}
}

func (ts *TutorStudentService) AddStudent(tutorID, studentID int64) error {
	err := ts.storage.AddStudent(tutorID, studentID)
	if err != nil {
		return err
	}

	return nil
}

func (ts *TutorStudentService) DeleteOneStudent(tutorID, studentID int64) error {
	err := ts.storage.DeleteOneStudent(tutorID, studentID)
	if err != nil {
		return err
	}

	return nil
}

func (ts *TutorStudentService) DeleteStudents(tutorID int64, studentIDs []int64) error {
	err := ts.storage.DeleteStudents(tutorID, studentIDs)
	if err != nil {
		return err
	}

	return nil
}

func (ts *TutorStudentService) GetTutorStudents(
	getTutorStudents dto.GetTutorStudents,
) ([]models.TutorStudent, int64, error) {
	res, count, err := ts.storage.GetTutorStudents(getTutorStudents)
	if err != nil {
		return nil, 0, err
	}

	return res, count, nil
}

func (ts *TutorStudentService) GetOneTutorStudentsCount(tutorID int64) (int64, error) {
	res, err := ts.storage.GetOneTutorStudentsCount(tutorID, nil)
	if err != nil {
		return 0, err
	}

	return res, nil
}

func (ts *TutorStudentService) GetTutorsStudentsCount(tutorIDs []int64) ([]int64, error) {
	res, err := ts.storage.GetTutorsStudentsCount(tutorIDs)
	if err != nil {
		return nil, err
	}

	return res, nil
}
