package groupDto

type UpdateGroupRequest struct {
	Title       string
	Description string
	TgGroupLink *string
	TgChatId    *int64
}
