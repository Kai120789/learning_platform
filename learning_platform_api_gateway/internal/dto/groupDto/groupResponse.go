package groupDto

import (
	"learning-platform/api-gateway/internal/dto/subjectDto"
	"learning-platform/api-gateway/internal/dto/userDto"
)

type GroupResponse struct {
	ID          int64   `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	SubjectID   int64   `json:"subject_id"`
	TutorID     int64   `json:"tutor_id"`
	TgGroupLink *string `json:"tg_group_link"`
	TgChatID    *string `json:"tg_chat_id"`
}

type GroupFullResponse struct {
	ID          int64                   `json:"id"`
	Title       string                  `json:"title"`
	Description string                  `json:"description"`
	Subject     subjectDto.Subject      `json:"subject"`
	Users       []userDto.UserShortInfo `json:"users"`
	TutorID     int64                   `json:"tutor_id"`
	TgGroupLink *string                 `json:"tg_group_link"`
	TgChatID    *string                 `json:"tg_chat_id"`
}
