package models

import "github.com/jackc/pgx/v5/pgtype"

type MaterialUser struct {
	ID         int64              `json:"id"`
	MaterialID int64              `json:"material_id"`
	UserID     int64              `json:"user_id"`
	CreatedAt  pgtype.Timestamptz `json:"created_at"`
	UpdatedAt  pgtype.Timestamptz `json:"updated_at"`
}
