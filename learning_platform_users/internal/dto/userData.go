package dto

import "learning-platform/users/internal/models/enum"

type UserData struct {
	UserID       int64                `json:"user_id"`
	Email        string               `json:"email"`
	Role         enum.UserRole        `json:"role"`
	Status       enum.UserStatus      `json:"status"`
	UserInfo     UserInfoResponse     `json:"user_info"`
	UserSettings UserSettingsResponse `json:"user_settings"`
}
