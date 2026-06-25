package dto

type UpdateGroup struct {
	ID          int64   `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	TgGroupLink *string `json:"tg_group_link"`
	TgChatId    *string `json:"tg_chat_id"`
}
