package models

import "github.com/jackc/pgx/v5/pgtype"

type UserSettings struct {
	UserID                 int64              `json:"user_id"`
	Is2FaEnabled           bool               `json:"is_2fa_enabled"`
	IsNotificationsEnabled bool               `json:"is_notifications_enabled"`
	CreatedAt              pgtype.Timestamptz `json:"created_at"`
	UpdatedAt              pgtype.Timestamptz `json:"updated_at"`
}
