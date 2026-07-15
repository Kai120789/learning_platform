package models

import (
	"github.com/jackc/pgx/v5/pgtype"
	"learning-platform/users/internal/models/enum"
)

type UserInfo struct {
	UserID     int64              `json:"user_id"`
	Name       string             `json:"name"`
	Surname    string             `json:"surname"`
	Patronymic pgtype.Text        `json:"patronymic"`
	City       pgtype.Text        `json:"city"`
	About      pgtype.Text        `json:"about"`
	Avatar     pgtype.Text        `json:"avatar"`
	Gender     enum.UserGender    `json:"gender"`
	BirthDate  pgtype.Date        `json:"birth_date"`
	CreatedAt  pgtype.Timestamptz `json:"created_at"`
	UpdatedAt  pgtype.Timestamptz `json:"updated_at"`
}
