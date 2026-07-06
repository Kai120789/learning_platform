package subjectDto

type UpdateUserSubjects struct {
	SubjectIDs        []int64 `json:"subject_ids"`
	DeletedSubjectIDs []int64 `json:"deleted_subject_ids"`
}
