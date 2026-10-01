package postgresModels

import "github.com/jackc/pgx/v5/pgtype"

type Board struct {
	ID        int64              `json:"id"`
	Title     string             `json:"title"`
	TutorID   int64              `json:"tutor_id"`
	CreatedAt pgtype.Timestamptz `json:"created_at"`
	UpdatedAt pgtype.Timestamptz `json:"updated_at"`
}
