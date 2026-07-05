package userDto

import "learning-platform/api-gateway/internal/dto/enum"

type UserData struct {
	UserID       int64           `json:"user_id"`
	Email        string          `json:"email"`
	Role         enum.UserRole   `json:"role"`
	Status       enum.UserStatus `json:"status"`
	UserInfo     UserInfo        `json:"user_info"`
	UserSettings UserSettings    `json:"user_settings"`
}
