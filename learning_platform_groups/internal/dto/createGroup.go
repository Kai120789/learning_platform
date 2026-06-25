package dto

type CreateGroup struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	SubjectID   int64   `json:"subject_id"`
	TutorID     int64   `json:"tutor_id"`
	TgGroupLink *string `json:"tg_group_link"`
	TgChatID    *string `json:"tg_chat_id"`
}
