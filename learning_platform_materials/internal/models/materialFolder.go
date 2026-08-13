package models

import "github.com/jackc/pgx/v5/pgtype"

type MaterialFolder struct {
	ID             int64              `json:"id"`
	Title          string             `json:"title"`
	ParentFolderID pgtype.Int8        `json:"parent_folder_id"`
	TutorID        int64              `json:"tutor_id"`
	CreatedAt      pgtype.Timestamptz `json:"created_at"`
	UpdatedAt      pgtype.Timestamptz `json:"updated_at"`
}
