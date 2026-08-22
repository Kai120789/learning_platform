package dto

type GetTutorStudents struct {
	TutorID              int64  `json:"tutor_id"`
	InteractedWithinDays *int64 `json:"interacted_within_days"`
	Page                 int64  `json:"page"`
	Limit                int64  `json:"limit"`
}
