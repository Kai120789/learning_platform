package models

import "github.com/jackc/pgx/v5/pgtype"

type UserGroup struct {
	ID        int64              `json:"id"`
	UserID    int64              `json:"user_id"`
	GroupID   int64              `json:"group_id"`
	CreatedAt pgtype.Timestamptz `json:"created_at"`
	UpdatedAt pgtype.Timestamptz `json:"updated_at"`
}
