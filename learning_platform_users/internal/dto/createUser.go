package dto

import "learning-platform/users/internal/models/enum"

type CreateUser struct {
	Email        string        `json:"email"`
	Name         string        `json:"name"`
	Surname      string        `json:"surname"`
	LastName     string        `json:"lastName"`
	Role         enum.UserRole `json:"role"`
	PasswordHash string        `json:"password_hash"`
}
