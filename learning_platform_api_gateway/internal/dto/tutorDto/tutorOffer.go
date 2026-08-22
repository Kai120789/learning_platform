package tutorDto

import "learning-platform/api-gateway/internal/dto/subjectDto"

type TutorOffer struct {
	ID              int64   `json:"id"`
	TutorID         int64   `json:"tutor_id"`
	SubjectID       int64   `json:"subject_id"`
	Title           string  `json:"title"`
	Description     *string `json:"description"`
	Price           int64   `json:"price"`
	DurationMinutes *int64  `json:"duration_minutes"`
}

type TutorOfferWithInfo struct {
	ID              int64              `json:"id"`
	TutorID         int64              `json:"tutor_id"`
	Subject         subjectDto.Subject `json:"subject"`
	Title           string             `json:"title"`
	Description     *string            `json:"description"`
	Price           int64              `json:"price"`
	DurationMinutes *int64             `json:"duration_minutes"`
}
