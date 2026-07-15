package dto

type LoginRequest struct {
	UserID   int64  `json:"user_id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	SessionID string `json:"session_id"`
}
