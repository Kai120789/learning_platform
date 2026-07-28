package dto

import "learning-platform/auth/internal/dto/enum"

type RegisterRequest struct {
	UserID int64         `json:"user_id"`
	Email  string        `json:"email"`
	Role   enum.UserRole `json:"role"`
}

type RegisterResponse struct {
	SessionID string
}
