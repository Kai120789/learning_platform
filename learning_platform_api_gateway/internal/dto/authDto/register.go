package authDto

import "learning-platform/api-gateway/internal/dto/enum"

type RegisterRequest struct {
	Email    string        `json:"email"`
	Name     string        `json:"name"`
	Surname  string        `json:"surname"`
	LastName *string       `json:"lastname"`
	Role     enum.UserRole `json:"role"`
	Password string        `json:"password"`
}

type RegisterResponse struct {
	SessionID string `json:"session_id"`
	UserID    int64  `json:"user_id"`
}
