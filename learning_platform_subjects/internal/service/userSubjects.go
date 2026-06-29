package service

import "learning-platform/subjects/internal/models"

type UserSubjectService struct {
	storage UserSubjectStorage
}

type UserSubjectStorage interface {
	GetAllUserSubjects(userID int64) ([]models.Subject, error)
	SetUserSubjects(userID int64, subjectIDs []int64) error
	DeleteUserSubjects(userID int64, deletedSubjectIDs []int64) error
}

func NewUserSubjectService(storage UserSubjectStorage) *UserSubjectService {
	return &UserSubjectService{
		storage: storage,
	}
}

func (us *UserSubjectService) GetUserSubjects(userID int64) ([]models.Subject, error) {
	res, err := us.storage.GetAllUserSubjects(userID)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (us *UserSubjectService) SetUserSubjects(userID int64, subjectIDs []int64) ([]models.Subject, error) {
	err := us.storage.SetUserSubjects(userID, subjectIDs)
	if err != nil {
		return nil, err
	}

	res, err := us.storage.GetAllUserSubjects(userID)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (us *UserSubjectService) UpdateUserSubjects(
	userID int64,
	subjectIDs []int64,
	deletedSubjectIDs []int64,
) ([]models.Subject, error) {
	err := us.storage.DeleteUserSubjects(userID, deletedSubjectIDs)
	if err != nil {
		return nil, err
	}

	err = us.storage.SetUserSubjects(userID, subjectIDs)
	if err != nil {
		return nil, err
	}

	res, err := us.storage.GetAllUserSubjects(userID)
	if err != nil {
		return nil, err
	}

	return res, nil
}
