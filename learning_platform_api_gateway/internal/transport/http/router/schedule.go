package router

import (
	"github.com/go-chi/chi/v5"
	"learning-platform/api-gateway/internal/dto/enum"
	"net/http"
)

type ScheduleRouter struct{}

type ScheduleHandler interface {
	GetAllSchedules(w http.ResponseWriter, r *http.Request)
	GetScheduleByID(w http.ResponseWriter, r *http.Request)
	GetSchedulesByTutorID(w http.ResponseWriter, r *http.Request)
	CreateSchedule(w http.ResponseWriter, r *http.Request)
	UpdateSchedule(w http.ResponseWriter, r *http.Request)
	DeleteSchedule(w http.ResponseWriter, r *http.Request)
	CreateScheduleSlot(w http.ResponseWriter, r *http.Request)
	UpdateScheduleSlot(w http.ResponseWriter, r *http.Request)
	DeleteScheduleSlot(w http.ResponseWriter, r *http.Request)
	BindLessonToScheduleSlot(w http.ResponseWriter, r *http.Request)
	DeleteLessonFromScheduleSlot(w http.ResponseWriter, r *http.Request)
}

func NewScheduleRouter() *ScheduleRouter {
	return &ScheduleRouter{}
}

func (s *ScheduleRouter) ScheduleRoutes(
	r chi.Router,
	h ScheduleHandler,
	jwtMiddleware func(http.Handler) http.Handler,
	roleMiddleware func(minNeededRole enum.UserRole) func(http.Handler) http.Handler,
) {
	r.With(jwtMiddleware).Route("/api/schedule", func(r chi.Router) {
		r.With(roleMiddleware(enum.RoleStudent)).Get("/", h.GetAllSchedules) // TODO: удалить потом, пока для тестов, на клиенте не нужно
		r.With(roleMiddleware(enum.RoleTutor)).Get("/tutor/{tutorID}", h.GetSchedulesByTutorID)
		r.With(roleMiddleware(enum.RoleTutor)).Get("/{scheduleID}", h.GetScheduleByID)
		r.With(roleMiddleware(enum.RoleTutor)).Post("/", h.CreateSchedule)
		r.With(roleMiddleware(enum.RoleTutor)).Put("/{scheduleID}", h.UpdateSchedule)
		r.With(roleMiddleware(enum.RoleTutor)).Delete("/{scheduleID}", h.DeleteSchedule)
		r.With(roleMiddleware(enum.RoleTutor)).Post("/slot/", h.CreateScheduleSlot)
		r.With(roleMiddleware(enum.RoleTutor)).Put("/slot/{scheduleSlotID}", h.UpdateScheduleSlot)
		r.With(roleMiddleware(enum.RoleTutor)).Delete("/slot/{scheduleSlotID}", h.DeleteScheduleSlot)
		r.With(roleMiddleware(enum.RoleTutor)).Patch("/slot/{scheduleSlotID}", h.BindLessonToScheduleSlot)
		r.With(roleMiddleware(enum.RoleTutor)).Delete("/slot/{scheduleSlotID}/lesson", h.DeleteLessonFromScheduleSlot)
	})
}
