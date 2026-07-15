package models

import (
	"github.com/jackc/pgx/v5/pgtype"
	"learning-platform/users/internal/models/enum"
)

type UserSettings struct {
	UserID                 int64              `json:"user_id"`
	Is2FaEnabled           bool               `json:"is_2fa_enabled"`
	IsNotificationsEnabled bool               `json:"is_notifications_enabled"`
	Language               enum.UserLanguage  `json:"language"`
	Theme                  enum.UserTheme     `json:"theme"`
	CreatedAt              pgtype.Timestamptz `json:"created_at"`
	UpdatedAt              pgtype.Timestamptz `json:"updated_at"`
}
