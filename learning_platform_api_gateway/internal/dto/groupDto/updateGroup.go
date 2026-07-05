package groupDto

type UpdateGroupRequest struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	TgGroupLink *string `json:"tg_group_link"`
	TgChatID    *string `json:"tg_chat_id"`
}
