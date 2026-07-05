package userDto

type UserSettings struct {
	UserID                 int64 `json:"user_id"`
	Is2FaEnabled           bool  `json:"is_2_fa_enabled"`
	IsNotificationsEnabled bool  `json:"is_notifications_enabled"`
}
