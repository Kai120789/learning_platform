package models

import "github.com/jackc/pgx/v5/pgtype"

type TutorReview struct {
	ID        int64              `json:"id"`
	TutorID   int64              `json:"tutor_id"`
	AuthorID  int64              `json:"author_id"`
	SubjectID int64              `json:"subject_id"`
	Text      string             `json:"text"`
	Rating    int64              `json:"rating"`
	CreatedAt pgtype.Timestamptz `json:"created_at"`
	UpdatedAt pgtype.Timestamptz `json:"updated_at"`
}
