package dto

import "learning-platform/users/internal/models/enum"

type UserSettingsRequest struct {
	UserID                 int64             `json:"user_id"`
	Is2FaEnabled           bool              `json:"is_2_fa_enabled"`
	IsNotificationsEnabled bool              `json:"is_notifications_enabled"`
	Language               enum.UserLanguage `json:"language"`
}

type UserSettingsResponse struct {
	Is2FaEnabled           bool              `json:"is_2_fa_enabled"`
	IsNotificationsEnabled bool              `json:"is_notifications_enabled"`
	Language               enum.UserLanguage `json:"language"`
	Theme                  enum.UserTheme    `json:"theme"`
}
