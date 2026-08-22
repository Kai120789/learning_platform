package models

import "github.com/jackc/pgx/v5/pgtype"

type TutorSubject struct {
	ID        int64              `json:"id"`
	TutorID   int64              `json:"tutor_id"`
	SubjectID int64              `json:"subject_id"`
	CreatedAt pgtype.Timestamptz `json:"created_at"`
	UpdatedAt pgtype.Timestamptz `json:"updated_at"`
}
