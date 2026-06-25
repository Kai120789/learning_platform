package models

import "github.com/jackc/pgx/v5/pgtype"

type UserInfo struct {
	UserID    int64              `json:"user_id"`
	Name      string             `json:"name"`
	Surname   string             `json:"surname"`
	Lastname  pgtype.Text        `json:"lastname"`
	City      pgtype.Text        `json:"city"`
	About     pgtype.Text        `json:"about"`
	CreatedAt pgtype.Timestamptz `json:"created_at"`
	UpdatedAt pgtype.Timestamptz `json:"updated_at"`
}
