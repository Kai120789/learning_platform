package dto

type RegisterRequest struct {
	UserID   int64  `json:"user_id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterResponse struct {
	SessionID string
}
