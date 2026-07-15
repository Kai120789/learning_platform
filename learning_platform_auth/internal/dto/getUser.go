package dto

type GetUser struct {
	UserID       int64  `json:"user_id"`
	Email        string `json:"email"`
	PasswordHash string `json:"password_hash"`
}
