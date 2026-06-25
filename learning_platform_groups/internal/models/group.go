package models

type Group struct {
	Id          int64
	Title       string
	Description string
	SubjectId   int64
	TutorId     int64
	TgGroupLink string
	TgChatId    int64
}
