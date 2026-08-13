package models

import "github.com/jackc/pgx/v5/pgtype"

type MaterialFolderUser struct {
	ID        int64              `json:"id"`
	FolderID  int64              `json:"folder_id"`
	UserID    int64              `json:"user_id"`
	CreatedAt pgtype.Timestamptz `json:"created_at"`
	UpdatedAt pgtype.Timestamptz `json:"updated_at"`
}
