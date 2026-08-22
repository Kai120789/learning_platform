package models

import "github.com/jackc/pgx/v5/pgtype"

type TutorStudent struct {
	ID               int64              `json:"id"`
	TutorID          int64              `json:"tutor_id"`
	StudentID        int64              `json:"student_id"`
	LastInteractedAt pgtype.Timestamptz `json:"last_interacted_at"`
	CreatedAt        pgtype.Timestamptz `json:"created_at"`
	UpdatedAt        pgtype.Timestamptz `json:"updated_at"`
}
