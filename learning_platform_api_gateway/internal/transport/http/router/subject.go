package router

import (
	"github.com/go-chi/chi/v5"
	"learning-platform/api-gateway/internal/dto/enum"
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
	roleMiddleware func(minNeededRole enum.UserRole) func(http.Handler) http.Handler,
) {
	r.With(jwtMiddleware).Route("/api/subject", func(r chi.Router) {
		r.With(roleMiddleware(enum.RoleStudent)).Get("/{subjectID}", h.GetOneSubject)
		r.With(roleMiddleware(enum.RoleStudent)).Get("/", h.GetAllSubjects)
		r.With(roleMiddleware(enum.RoleStudent)).Get("/user/{userID}", h.GetUserSubjects)
		r.With(roleMiddleware(enum.RoleStudent)).Post("/user/{userID}", h.SetUserSubjects)
		r.With(roleMiddleware(enum.RoleStudent)).Put("/user/{userID}", h.UpdateUserSubjects)
	})
}
