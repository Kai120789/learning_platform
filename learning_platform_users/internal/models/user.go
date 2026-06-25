package models

import (
	"github.com/jackc/pgx/v5/pgtype"
	"learning-platform/users/internal/models/enum"
)

type User struct {
	ID           int64              `json:"id"`
	Email        string             `json:"email"`
	PasswordHash string             `json:"password_hash"`
	Role         enum.UserRole      `json:"role"`
	Status       enum.UserStatus    `json:"status"`
	CreatedAt    pgtype.Timestamptz `json:"created_at"`
	UpdatedAt    pgtype.Timestamptz `json:"updated_at"`
}
