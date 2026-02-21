package dto

type GetUser struct {
	UserId       int64
	Email        string
	PasswordHash string
}
