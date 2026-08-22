package dto

type GetTutors struct {
	SubjectID *int64 `json:"subject_id"`
	Page      int64  `json:"page"`
	Limit     int64  `json:"limit"`
}
