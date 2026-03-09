package dto

type LoginRequest struct {
	UserId   int64
	Email    string
	Password string
}

type LoginResponse struct {
	AccessToken string
	UserId      int64
}
