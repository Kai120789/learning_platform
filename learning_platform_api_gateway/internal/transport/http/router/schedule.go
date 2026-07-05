package router

import (
	"github.com/go-chi/chi/v5"
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
	UpdateScheduleSlot(w http.ResponseWriter, r *http.Request)
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
) {
	r.With(jwtMiddleware).Route("/api/schedule", func(r chi.Router) {
		r.Get("/", h.GetAllSchedules)
		r.Get("/{scheduleId}", h.GetScheduleByID)
		r.Get("/tutor/{scheduleId}", h.GetSchedulesByTutorID)
		r.Post("/", h.CreateSchedule)
		r.Put("/{scheduleId}", h.UpdateSchedule)
		r.Delete("/{scheduleId}", h.DeleteSchedule)
		r.Put("/slot/{scheduleSlotId}", h.UpdateScheduleSlot)
		r.Patch("/slot/{scheduleSlotId}", h.BindLessonToScheduleSlot)
		r.Delete("/slot/{scheduleSlotId}", h.DeleteLessonFromScheduleSlot)
	})
}
