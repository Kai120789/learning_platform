package handler

import (
	"go.uber.org/zap"
	"net/http"
)

type ScheduleHandler struct {
	service ScheduleService
	logger  *zap.Logger
}

type ScheduleService interface{}

func NewScheduleHandler(service ScheduleService, logger *zap.Logger) *ScheduleHandler {
	return &ScheduleHandler{
		service: service,
		logger:  logger,
	}
}

func (s *ScheduleHandler) GetAllSchedules(w http.ResponseWriter, r *http.Request) {

}

func (s *ScheduleHandler) GetScheduleByID(w http.ResponseWriter, r *http.Request) {

}

func (s *ScheduleHandler) GetSchedulesByTutorID(w http.ResponseWriter, r *http.Request) {

}

func (s *ScheduleHandler) CreateSchedule(w http.ResponseWriter, r *http.Request) {

}

func (s *ScheduleHandler) UpdateSchedule(w http.ResponseWriter, r *http.Request) {

}

func (s *ScheduleHandler) DeleteSchedule(w http.ResponseWriter, r *http.Request) {

}

func (s *ScheduleHandler) UpdateScheduleSlot(w http.ResponseWriter, r *http.Request) {

}

func (s *ScheduleHandler) BindLessonToScheduleSlot(w http.ResponseWriter, r *http.Request) {

}

func (s *ScheduleHandler) DeleteLessonFromScheduleSlot(w http.ResponseWriter, r *http.Request) {

}
