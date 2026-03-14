package dto

type LoginRequest struct {
	UserId   int64
	Email    string
	Password string
}

type LoginResponse struct {
	SessionId string
	UserId    int64
}
