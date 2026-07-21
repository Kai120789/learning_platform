package userDto

type UserShortInfo struct {
	ID         int64   `json:"id"`
	Name       string  `json:"name"`
	Surname    string  `json:"surname"`
	Patronymic *string `json:"patronymic"`
	TgUsername *string `json:"tg_username"`
}
