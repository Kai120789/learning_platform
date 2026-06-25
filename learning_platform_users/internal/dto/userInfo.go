package dto

type UserInfo struct {
	UserID   int64   `json:"user_id"`
	Name     string  `json:"name"`
	Surname  string  `json:"surname"`
	Lastname *string `json:"lastname"`
	City     *string `json:"city"`
	About    *string `json:"about"`
}
