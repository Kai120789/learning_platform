package tutorDto

type NewTutorOffer struct {
	TutorID         int64   `json:"tutor_id"`
	SubjectID       int64   `json:"subject_id"`
	Title           string  `json:"title"`
	Description     *string `json:"description"`
	Price           int64   `json:"price"`
	DurationMinutes *int64  `json:"duration_minutes"`
}
