package router

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

type SubjectRouter struct{}

type SubjectHandler interface {
	GetOneSubject(w http.ResponseWriter, r *http.Request)
	GetAllSubjects(w http.ResponseWriter, r *http.Request)
	GetUserSubjects(w http.ResponseWriter, r *http.Request)
	SetUserSubjects(w http.ResponseWriter, r *http.Request)
	UpdateUserSubjects(w http.ResponseWriter, r *http.Request)
}

func NewSubjectRouter() *SubjectRouter {
	return &SubjectRouter{}
}

func (s *SubjectRouter) SubjectRoutes(
	r chi.Router,
	h SubjectHandler,
	jwtMiddleware func(handler http.Handler) http.Handler,
) {
	r.With(jwtMiddleware).Route("/api/subject", func(r chi.Router) {
		r.Get("/{subjectId}", h.GetOneSubject)
		r.Get("/", h.GetAllSubjects)
		r.Get("/user/{userId}", h.GetUserSubjects)
		r.Post("/user/{userId}", h.SetUserSubjects)
		r.Put("/user/{userId}", h.UpdateUserSubjects)
	})
}
