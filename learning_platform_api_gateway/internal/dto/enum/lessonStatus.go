package enum

type LessonStatus string

const (
	StatusScheduled LessonStatus = "SCHEDULED"
	StatusInProcess LessonStatus = "IN_PROCESS"
	StatusCompleted LessonStatus = "COMPLETED"
	StatusCancelled LessonStatus = "CANCELLED"
)
