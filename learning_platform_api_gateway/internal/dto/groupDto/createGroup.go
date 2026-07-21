package groupDto

type CreateGroupRequest struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	SubjectID   int64   `json:"subject_id"`
	TgGroupLink *string `json:"tg_group_link"`
	TgChatID    *string `json:"tg_chat_id"`
}
