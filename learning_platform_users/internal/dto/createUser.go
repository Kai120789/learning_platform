package dto

import (
	"learning-platform/users/internal/models/enum"
	"time"
)

type CreateUser struct {
	Email        string            `json:"email"`
	Name         string            `json:"name"`
	Surname      string            `json:"surname"`
	Patronymic   *string           `json:"patronymic"`
	Role         enum.UserRole     `json:"role"`
	Gender       enum.UserGender   `json:"gender"`
	Language     enum.UserLanguage `json:"language"`
	PasswordHash string            `json:"password_hash"`
	BirthDate    *time.Time        `json:"birth_date"`
}
