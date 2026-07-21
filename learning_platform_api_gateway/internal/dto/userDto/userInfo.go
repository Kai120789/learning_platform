package userDto

import (
	"learning-platform/api-gateway/internal/dto/enum"
	"time"
)

type UserInfoRequest struct {
	Name       string          `json:"name"`
	Surname    string          `json:"surname"`
	Patronymic *string         `json:"patronymic"`
	City       *string         `json:"city"`
	About      *string         `json:"about"`
	Gender     enum.UserGender `json:"gender"`
	BirthDate  *time.Time      `json:"birth_date"`
}

type UserInfoResponse struct {
	Name       string          `json:"name"`
	Surname    string          `json:"surname"`
	Patronymic *string         `json:"patronymic"`
	TgUsername *string         `json:"tg_username"`
	City       *string         `json:"city"`
	About      *string         `json:"about"`
	Avatar     *string         `json:"avatar"`
	Gender     enum.UserGender `json:"gender"`
	BirthDate  *time.Time      `json:"birth_date"`
}
