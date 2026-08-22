package service

import (
	"learning-platform/tutors/internal/dto"
	"learning-platform/tutors/internal/models"
)

type TutorService struct {
	review  ReviewService
	offer   OfferService
	student StudentService
	subject SubjectService
}

func NewTutorService(
	review ReviewService,
	offer OfferService,
	student StudentService,
	subject SubjectService,
) *TutorService {
	return &TutorService{
		review:  review,
		offer:   offer,
		student: student,
		subject: subject,
	}
}

type ReviewService interface {
	GetOneTutorReviewsCount(tutorID int64) (int64, error)
	GetTutorsReviewsCount(tutorIDs []int64) ([]int64, error)
	GetReviewByTutorAndAuthor(tutorID, userID int64) (*models.TutorReview, error)
	GetOneTutorRating(tutorID int64) (float32, error)
	GetTutorsRatings(tutorIDs []int64) ([]float32, error)
}

type OfferService interface {
	GetTutorOffers(tutorID int64) ([]models.TutorOffer, error)
}

type SubjectService interface {
	GetOneTutorSubjectIDs(tutorID int64) ([]int64, error)
	GetTutorsSubjectIDs(tutorIDs []int64) ([][]int64, error)
	GetTutorsBySubject(getTutors dto.GetTutors) ([]int64, int64, error)
}

type StudentService interface {
	GetOneTutorStudentsCount(tutorID int64) (int64, error)
	GetTutorsStudentsCount(tutorIDs []int64) ([]int64, error)
}

func (t *TutorService) GetOneTutor(tutorID, userID int64) (*dto.OneTutor, error) {
	rating, err := t.review.GetOneTutorRating(tutorID)
	if err != nil {
		return nil, err
	}

	reviewsCount, err := t.review.GetOneTutorReviewsCount(tutorID)
	if err != nil {
		return nil, err
	}

	studentsCount, err := t.student.GetOneTutorStudentsCount(tutorID)
	if err != nil {
		return nil, err
	}

	offers, err := t.offer.GetTutorOffers(tutorID)
	if err != nil {
		return nil, err
	}

	subjectIDs, err := t.subject.GetOneTutorSubjectIDs(tutorID)
	if err != nil {
		return nil, err
	}

	myReview, err := t.review.GetReviewByTutorAndAuthor(tutorID, userID)
	if err != nil {
		return nil, err
	}

	return &dto.OneTutor{
		TutorInfo: dto.TutorShortInfo{
			TutorID:       tutorID,
			Rating:        rating,
			ReviewsCount:  reviewsCount,
			StudentsCount: studentsCount,
			SubjectIDs:    subjectIDs,
		},
		Offers:   offers,
		MyReview: myReview,
	}, nil
}

func (t *TutorService) GetTutors(getTutors dto.GetTutors) ([]dto.TutorShortInfo, int64, error) {
	tutorIDs, count, err := t.subject.GetTutorsBySubject(getTutors)
	if err != nil {
		return nil, 0, err
	}

	ratings, err := t.review.GetTutorsRatings(tutorIDs)
	if err != nil {
		return nil, 0, err
	}

	reviewsCounts, err := t.review.GetTutorsReviewsCount(tutorIDs)
	if err != nil {
		return nil, 0, err
	}

	subjectsIDs, err := t.subject.GetTutorsSubjectIDs(tutorIDs)
	if err != nil {
		return nil, 0, err
	}

	studentsCounts, err := t.student.GetTutorsStudentsCount(tutorIDs)
	if err != nil {
		return nil, 0, err
	}

	resTutors := make([]dto.TutorShortInfo, len(tutorIDs))
	for ind, oneTutorID := range tutorIDs {
		resTutors[ind] = dto.TutorShortInfo{
			TutorID:       oneTutorID,
			Rating:        ratings[ind],
			ReviewsCount:  reviewsCounts[ind],
			StudentsCount: studentsCounts[ind],
			SubjectIDs:    subjectsIDs[ind],
		}
	}

	return resTutors, count, nil
}
