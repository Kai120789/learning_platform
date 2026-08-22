package dto

type TutorShortInfo struct {
	TutorID       int64   `json:"tutor_id"`
	Rating        float32 `json:"rating"`
	ReviewsCount  int64   `json:"reviews_count"`
	StudentsCount int64   `json:"students_count"`
	SubjectIDs    []int64 `json:"subject_ids"`
}
