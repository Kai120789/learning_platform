package dto

import "learning-platform/auth/internal/dto/enum"

type LoginRequest struct {
	UserID int64         `json:"user_id"`
	Email  string        `json:"email"`
	Role   enum.UserRole `json:"role"`
}

type LoginResponse struct {
	SessionID string `json:"session_id"`
}
