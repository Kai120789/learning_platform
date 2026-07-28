package dto

import "learning-platform/auth/internal/dto/enum"

type CreateJWT struct {
	UserID      int64         `json:"user_id"`
	Email       string        `json:"email "`
	Role        enum.UserRole `json:"role"`
	SignedKey   string        `json:"signed_key"`
	SessionID   *string       `json:"session_id"`
	Issuer      string        `json:"issuer"`
	AccessTime  int64         `json:"access_time"`
	RefreshTime int64         `json:"refresh_time"`
}
