package tutorDto

import (
	"learning-platform/api-gateway/internal/dto/subjectDto"
	"learning-platform/api-gateway/internal/dto/userDto"
)

type TutorShortInfo struct {
	TutorID       int64   `json:"tutor_id"`
	Rating        float32 `json:"rating"`
	ReviewsCount  int64   `json:"reviews_count"`
	StudentsCount int64   `json:"students_count"`
	SubjectIDs    []int64 `json:"subject_ids"`
}

type TutorShortInfoWithSubjects struct {
	Tutor         userDto.UserShortInfo `json:"tutor"`
	Rating        float32               `json:"rating"`
	ReviewsCount  int64                 `json:"reviews_count"`
	StudentsCount int64                 `json:"students_count"`
	Subjects      []subjectDto.Subject  `json:"subjects"`
}
