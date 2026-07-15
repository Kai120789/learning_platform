package authDto

import (
	"learning-platform/api-gateway/internal/dto/enum"
	"time"
)

type RegisterRequest struct {
	Email      string            `json:"email"`
	Name       string            `json:"name"`
	Surname    string            `json:"surname"`
	Patronymic *string           `json:"patronymic"`
	Role       enum.UserRole     `json:"role"`
	Gender     enum.UserGender   `json:"gender"`
	Language   enum.UserLanguage `json:"language"`
	BirthDate  *time.Time        `json:"birth_date"`
	Password   string            `json:"password"`
}

type RegisterResponse struct {
	SessionID string `json:"session_id"`
}
