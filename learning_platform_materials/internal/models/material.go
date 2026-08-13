package models

import "github.com/jackc/pgx/v5/pgtype"

type Material struct {
	ID            int64              `json:"id"`
	Title         string             `json:"title"`
	Size          int64              `json:"size"`
	FolderID      pgtype.Int8        `json:"folder_id"`
	TutorID       int64              `json:"tutor_id"`
	MimeType      string             `json:"mime_type"`
	MediaObjectID pgtype.UUID        `json:"media_object_id"`
	CreatedAt     pgtype.Timestamptz `json:"created_at"`
	UpdatedAt     pgtype.Timestamptz `json:"updated_at"`
}
