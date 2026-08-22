package tutorDto

import (
	"learning-platform/api-gateway/internal/dto/subjectDto"
	"learning-platform/api-gateway/internal/dto/userDto"
	"time"
)

type TutorReview struct {
	ID        int64     `json:"id"`
	TutorID   int64     `json:"tutor_id"`
	AuthorID  int64     `json:"author_id"`
	SubjectID int64     `json:"subject_id"`
	Text      string    `json:"text"`
	Rating    int64     `json:"rating"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type TutorReviewWithInfo struct {
	ID        int64                 `json:"id"`
	Tutor     userDto.UserShortInfo `json:"tutor"`
	Author    userDto.UserShortInfo `json:"author"`
	Subject   subjectDto.Subject    `json:"subject"`
	Text      string                `json:"text"`
	Rating    int64                 `json:"rating"`
	CreatedAt time.Time             `json:"created_at"`
	UpdatedAt time.Time             `json:"updated_at"`
}
