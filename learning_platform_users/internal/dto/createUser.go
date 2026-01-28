package dto

type CreateUser struct {
	Email        string
	Name         string
	Surname      string
	LastName     string
	Role         string
	PasswordHash string
}
