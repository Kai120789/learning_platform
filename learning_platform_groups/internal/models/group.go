package models

import "github.com/jackc/pgx/v5/pgtype"

type Group struct {
	ID          int64              `json:"id"`
	Title       string             `json:"title"`
	Description string             `json:"description"`
	SubjectID   int64              `json:"subject_id"`
	TutorID     int64              `json:"tutor_id"`
	TgGroupLink pgtype.Text        `json:"tg_group_link"`
	TgChatID    pgtype.Text        `json:"tg_chat_id"`
	CreatedAt   pgtype.Timestamptz `json:"created_at"`
	UpdatedAt   pgtype.Timestamptz `json:"updated_at"`
}
