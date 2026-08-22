package tutorDto

import (
	"learning-platform/api-gateway/internal/dto/userDto"
	"time"
)

type TutorStudent struct {
	StudentID        int64      `json:"student_id"`
	LastInteractedAt *time.Time `json:"last_interacted_at"`
}

type TutorStudentWithInfo struct {
	Student          userDto.UserShortInfo `json:"student"`
	LastInteractedAt *time.Time            `json:"last_interacted_at"`
}
