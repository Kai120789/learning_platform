package userDto

import "learning-platform/api-gateway/internal/dto/enum"

type GetWithPagination struct {
	Search string        `json:"search"`
	Page   int64         `json:"page"`
	Limit  int64         `json:"limit"`
	Role   enum.UserRole `json:"role"`
}
