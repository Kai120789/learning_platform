package models

import (
	"github.com/jackc/pgx/v5/pgtype"
	"learning-platform/lessons/internal/models/enum"
)

type Lesson struct {
	ID        int64              `json:"id"`
	BoardID   pgtype.Int8        `json:"board_id"`
	MeetLink  pgtype.Text        `json:"meet_link"`
	StartTime pgtype.Timestamptz `json:"start_time"`
	Duration  int64              `json:"duration"`
	TutorID   int64              `json:"tutor_id"`
	Status    enum.LessonStatus  `json:"status"`
	CreatedAt pgtype.Timestamptz `json:"created_at"`
	UpdatedAt pgtype.Timestamptz `json:"updated_at"`
}
