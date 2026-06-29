package models

import (
	"github.com/jackc/pgx/v5/pgtype"
	"learning-platform/subjects/internal/models/enum"
)

type Subject struct {
	ID        int64              `json:"id"`
	Code      string             `json:"code"`
	Title     string             `json:"title"`
	Type      enum.Type          `json:"type"`
	CreatedAt pgtype.Timestamptz `json:"created_at"`
	UpdatedAt pgtype.Timestamptz `json:"updated_at"`
}
