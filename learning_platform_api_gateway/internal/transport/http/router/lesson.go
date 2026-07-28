package router

import (
	"github.com/go-chi/chi/v5"
	"learning-platform/api-gateway/internal/dto/enum"
	"net/http"
)

type LessonRouter struct{}

type LessonHandler interface {
	GetOneLesson(w http.ResponseWriter, r *http.Request)
	GetLessonsByUserID(w http.ResponseWriter, r *http.Request)
	CreateLesson(w http.ResponseWriter, r *http.Request)
	UpdateLesson(w http.ResponseWriter, r *http.Request)
	UpdateLessonStatus(w http.ResponseWriter, r *http.Request)
	GetLessonsByTutorID(w http.ResponseWriter, r *http.Request)
}

func NewLessonRouter() *LessonRouter {
	return &LessonRouter{}
}

func (l *LessonRouter) LessonRoutes(
	r chi.Router,
	h LessonHandler,
	jwtMiddleware func(http.Handler) http.Handler,
	roleMiddleware func(minNeededRole enum.UserRole) func(http.Handler) http.Handler,
) {
	r.With(jwtMiddleware).Route("/api/lesson", func(r chi.Router) {
		r.With(roleMiddleware(enum.RoleStudent)).Get("/{lessonID}", h.GetOneLesson)
		r.With(roleMiddleware(enum.RoleStudent)).Get("/user/{userID}", h.GetLessonsByUserID)
		r.With(roleMiddleware(enum.RoleTutor)).Get("/tutor/{tutorID}", h.GetLessonsByTutorID)
		r.With(roleMiddleware(enum.RoleTutor)).Post("/", h.CreateLesson)
		r.With(roleMiddleware(enum.RoleTutor)).Put("/{lessonID}", h.UpdateLesson)
		r.With(roleMiddleware(enum.RoleTutor)).Patch("/{lessonID}", h.UpdateLessonStatus)
	})
}
