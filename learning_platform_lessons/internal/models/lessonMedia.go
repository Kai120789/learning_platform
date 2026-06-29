package models

import (
	"github.com/jackc/pgx/v5/pgtype"
	"learning-platform/lessons/internal/models/enum"
)

type LessonMedia struct {
	ID        int64              `json:"id"`
	LessonID  int64              `json:"lesson_id"`
	S3Link    string             `json:"s3_link"`
	S3Preview string             `json:"s3_preview"`
	Type      enum.MediaType     `json:"type"`
	CreatedAt pgtype.Timestamptz `json:"created_at"`
	UpdatedAt pgtype.Timestamptz `json:"updated_at"`
}
