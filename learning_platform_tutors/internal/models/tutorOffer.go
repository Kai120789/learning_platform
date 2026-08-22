package models

import "github.com/jackc/pgx/v5/pgtype"

type TutorOffer struct {
	ID              int64              `json:"id"`
	TutorID         int64              `json:"tutor_id"`
	SubjectID       int64              `json:"subject_id"`
	Title           string             `json:"title"`
	Description     pgtype.Text        `json:"description"`
	Price           int64              `json:"price"`
	DurationMinutes pgtype.Int8        `json:"duration_minutes"`
	CreatedAt       pgtype.Timestamptz `json:"created_at"`
	UpdatedAt       pgtype.Timestamptz `json:"updated_at"`
}
