package dto

type NewTutorReview struct {
	Text      string `json:"text"`
	TutorID   int64  `json:"tutor_id"`
	AuthorID  int64  `json:"author_id"`
	SubjectID int64  `json:"subject_id"`
	Rating    int64  `json:"rating"`
}
