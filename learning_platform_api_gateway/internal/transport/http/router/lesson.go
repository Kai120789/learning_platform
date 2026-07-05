package router

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

type LessonRouter struct{}

type LessonHandler interface {
	GetOneLesson(w http.ResponseWriter, r *http.Request)
	GetLessonsByUserId(w http.ResponseWriter, r *http.Request)
	CreateLesson(w http.ResponseWriter, r *http.Request)
	UpdateLesson(w http.ResponseWriter, r *http.Request)
	UpdateLessonStatus(w http.ResponseWriter, r *http.Request)
	GetLessonsByTutorId(w http.ResponseWriter, r *http.Request)
}

func NewLessonRouter() *LessonRouter {
	return &LessonRouter{}
}

func (l *LessonRouter) LessonRoutes(
	r chi.Router,
	h LessonHandler,
	jwtMiddleware func(http.Handler) http.Handler,
) {
	r.With(jwtMiddleware).Route("/api/lesson", func(r chi.Router) {
		r.Get("/{lessonId}", h.GetOneLesson)
		r.Get("/user/{userId}", h.GetLessonsByUserId)
		r.Get("/tutor/{tutorId}", h.GetLessonsByTutorId)
		r.Post("/", h.CreateLesson)
		r.Put("/{lessonId}", h.UpdateLesson)
		r.Patch("/{lessonId}", h.UpdateLessonStatus)
	})
}
