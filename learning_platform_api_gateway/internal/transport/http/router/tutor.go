package router

import (
	"github.com/go-chi/chi/v5"
	"learning-platform/api-gateway/internal/dto/enum"
	"net/http"
)

type TutorRouter struct{}

type TutorHandler interface {
	AddStudent(w http.ResponseWriter, r *http.Request)
	DeleteOneStudent(w http.ResponseWriter, r *http.Request)
	DeleteStudents(w http.ResponseWriter, r *http.Request)
	GetTutorStudents(w http.ResponseWriter, r *http.Request)
	GetOneTutor(w http.ResponseWriter, r *http.Request)
	GetTutors(w http.ResponseWriter, r *http.Request)
	GetTutorReviews(w http.ResponseWriter, r *http.Request)
	AddTutorReview(w http.ResponseWriter, r *http.Request)
	UpdateTutorReview(w http.ResponseWriter, r *http.Request)
	DeleteTutorReview(w http.ResponseWriter, r *http.Request)
	GetTutorOffers(w http.ResponseWriter, r *http.Request)
	AddTutorOffer(w http.ResponseWriter, r *http.Request)
	UpdateTutorOffer(w http.ResponseWriter, r *http.Request)
	DeleteOneTutorOffer(w http.ResponseWriter, r *http.Request)
	DeleteTutorOffers(w http.ResponseWriter, r *http.Request)
	AddTutorSubjects(w http.ResponseWriter, r *http.Request)
	UpdateTutorSubjects(w http.ResponseWriter, r *http.Request)
}

func NewTutorRouter() *TutorRouter {
	return &TutorRouter{}
}

func (t *TutorRouter) TutorRoutes(
	r chi.Router,
	h TutorHandler,
	jwtMiddleware func(handler http.Handler) http.Handler,
	roleMiddleware func(minNeededRole enum.UserRole) func(http.Handler) http.Handler,
) {
	r.With(jwtMiddleware).Route("/api/tutor", func(r chi.Router) {
		r.With(roleMiddleware(enum.RoleStudent)).Get("/{tutorID}", h.GetOneTutor)
		r.With(roleMiddleware(enum.RoleStudent)).Get("/", h.GetTutors)
		r.With(roleMiddleware(enum.RoleTutor)).Route("/student", func(r chi.Router) {
			r.Post("/{studentID}", h.AddStudent)
			r.Delete("/{studentID}", h.DeleteOneStudent)
			r.Delete("/students", h.DeleteStudents)
			r.Get("/", h.GetTutorStudents)
		})
		r.With(roleMiddleware(enum.RoleStudent)).Route("/review", func(r chi.Router) {
			r.Get("/{tutorID}", h.GetTutorReviews)
			r.Post("/", h.AddTutorReview)
			r.Put("/", h.UpdateTutorReview)
			r.Delete("/{reviewID}", h.DeleteTutorReview)
		})
		r.Route("/offer", func(r chi.Router) {
			r.With(roleMiddleware(enum.RoleStudent)).Get("/{tutorID}", h.GetTutorOffers)
			r.With(roleMiddleware(enum.RoleTutor)).Post("/", h.AddTutorOffer)
			r.With(roleMiddleware(enum.RoleTutor)).Put("/", h.UpdateTutorOffer)
			r.With(roleMiddleware(enum.RoleTutor)).Delete("/{offerID}", h.DeleteOneTutorOffer)
			r.With(roleMiddleware(enum.RoleTutor)).Delete("/", h.DeleteTutorOffers)
		})
		r.With(roleMiddleware(enum.RoleTutor)).Route("/subject", func(r chi.Router) {
			r.Post("/", h.AddTutorSubjects)
			r.Put("/", h.UpdateTutorSubjects)
		})
	})
}
