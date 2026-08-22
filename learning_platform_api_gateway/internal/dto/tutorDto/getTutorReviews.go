package tutorDto

type GetTutorReviews struct {
	TutorID int64 `json:"tutor_id"`
	Page    int64 `json:"page"`
	Limit   int64 `json:"limit"`
}
