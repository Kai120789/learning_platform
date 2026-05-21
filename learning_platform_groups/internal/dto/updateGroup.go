package dto

type UpdateGroup struct {
	Id          int64
	Title       string
	Description string
	TgGroupLink *string
	TgChatId    *int64
}
