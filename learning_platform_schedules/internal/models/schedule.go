package models

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Schedule struct {
	ID        int64              `json:"id"`
	TutorID   int64              `json:"tutor_id"`
	StartTime pgtype.Timestamptz `json:"start_time"`
	EndTime   pgtype.Timestamptz `json:"end_time"`
	CreatedAt pgtype.Timestamptz `json:"created_at"`
	UpdatedAt pgtype.Timestamptz `json:"updated_at"`
}
