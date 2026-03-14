package dto

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	SessionId string `json:"session_id"`
	UserId    int64  `json:"user_id"`
}
