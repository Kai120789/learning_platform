package userDto

import "learning-platform/api-gateway/internal/dto/enum"

type GetUser struct {
	UserID       int64           `json:"user_id"`
	Email        string          `json:"email"`
	PasswordHash string          `json:"password_hash"`
	Role         enum.UserRole   `json:"role"`
	Status       enum.UserStatus `json:"status"`
}
