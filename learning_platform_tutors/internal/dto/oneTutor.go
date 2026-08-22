package dto

import "learning-platform/tutors/internal/models"

type OneTutor struct {
	TutorInfo TutorShortInfo      `json:"tutor_info"`
	Offers    []models.TutorOffer `json:"offers"`
	MyReview  *models.TutorReview `json:"my_review"`
}
