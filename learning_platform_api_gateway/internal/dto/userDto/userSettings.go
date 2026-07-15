package userDto

import "learning-platform/api-gateway/internal/dto/enum"

type UserSettings struct {
	Is2FaEnabled           bool              `json:"is_2_fa_enabled"`
	IsNotificationsEnabled bool              `json:"is_notifications_enabled"`
	Language               enum.UserLanguage `json:"language"`
	Theme                  enum.UserTheme    `json:"theme"`
}
