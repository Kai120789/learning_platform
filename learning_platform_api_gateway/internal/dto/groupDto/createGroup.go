package groupDto

type CreateGroupRequest struct {
	Title       string
	Description string
	SubjectId   int64
	TutorId     int64
	TgGroupLink *string
	TgChatId    *int64
}
