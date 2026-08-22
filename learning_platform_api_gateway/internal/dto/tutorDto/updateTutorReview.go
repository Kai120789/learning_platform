package tutorDto

type UpdateTutorReview struct {
	ID        int64  `json:"id"`
	Text      string `json:"text"`
	AuthorID  int64  `json:"author_id"`
	SubjectID int64  `json:"subject_id"`
	Rating    int64  `json:"rating"`
}
