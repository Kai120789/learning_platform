package service

import (
	"learning-platform/tutors/internal/dto"
	"learning-platform/tutors/internal/models"
)

type TutorReviewService struct {
	storage TutorReviewStorage
}

type TutorReviewStorage interface {
	GetTutorReviews(getTutorReviews dto.GetTutorReviews) ([]models.TutorReview, int64, error)
	AddTutorReview(review dto.NewTutorReview) (*models.TutorReview, error)
	UpdateTutorReview(updReview dto.UpdateTutorReview) (*models.TutorReview, error)
	DeleteTutorReview(reviewID, authorID int64) error
	ChackCanAddReview(tutorID, authorID int64) (bool, error)
	GetOneTutorReviewsCount(tutorID int64) (int64, error)
	GetTutorsReviewsCount(tutorIDs []int64) ([]int64, error)
	GetReviewByTutorAndAuthor(tutorID, userID int64) (*models.TutorReview, error)
	GetOneTutorRating(tutorID int64) (float32, error)
	GetTutorsRatings(tutorIDs []int64) ([]float32, error)
}

func NewTutorReviewService(storage TutorReviewStorage) *TutorReviewService {
	return &TutorReviewService{
		storage: storage,
	}
}

func (tr *TutorReviewService) GetTutorReviews(
	getTutorReviews dto.GetTutorReviews,
) ([]models.TutorReview, int64, error) {
	res, count, err := tr.storage.GetTutorReviews(getTutorReviews)
	if err != nil {
		return nil, 0, err
	}

	return res, count, nil
}

func (tr *TutorReviewService) AddTutorReview(review dto.NewTutorReview) (*models.TutorReview, error) {
	res, err := tr.storage.AddTutorReview(review)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (tr *TutorReviewService) UpdateTutorReview(updReview dto.UpdateTutorReview) (*models.TutorReview, error) {
	res, err := tr.storage.UpdateTutorReview(updReview)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (tr *TutorReviewService) DeleteTutorReview(reviewID, authorID int64) error {
	err := tr.storage.DeleteTutorReview(reviewID, authorID)
	if err != nil {
		return err
	}

	return nil
}

func (tr *TutorReviewService) ChackCanAddReview(tutorID, authorID int64) (bool, error) {
	canAdd, err := tr.storage.ChackCanAddReview(tutorID, authorID)
	if err != nil {
		return false, err
	}

	return canAdd, nil
}

func (tr *TutorReviewService) GetOneTutorReviewsCount(tutorID int64) (int64, error) {
	count, err := tr.storage.GetOneTutorReviewsCount(tutorID)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (tr *TutorReviewService) GetTutorsReviewsCount(tutorIDs []int64) ([]int64, error) {
	res, err := tr.storage.GetTutorsReviewsCount(tutorIDs)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (tr *TutorReviewService) GetReviewByTutorAndAuthor(tutorID, userID int64) (*models.TutorReview, error) {
	res, err := tr.storage.GetReviewByTutorAndAuthor(tutorID, userID)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (tr *TutorReviewService) GetOneTutorRating(tutorID int64) (float32, error) {
	res, err := tr.storage.GetOneTutorRating(tutorID)
	if err != nil {
		return 0, err
	}

	return res, nil
}

func (tr *TutorReviewService) GetTutorsRatings(tutorIDs []int64) ([]float32, error) {
	res, err := tr.storage.GetTutorsRatings(tutorIDs)
	if err != nil {
		return nil, err
	}

	return res, nil
}
