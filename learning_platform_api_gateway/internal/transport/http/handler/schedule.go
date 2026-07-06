package handler

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
	"learning-platform/api-gateway/internal/dto/scheduleDto"
	"net/http"
	"strconv"
)

type ScheduleHandler struct {
	service ScheduleService
	logger  *zap.Logger
}

type ScheduleService interface {
	GetAllSchedules() ([]scheduleDto.ScheduleResponse, error)
	GetScheduleByID(scheduleID int64) (*scheduleDto.ScheduleResponse, error)
	GetSchedulesByTutorID(tutorID int64) ([]scheduleDto.ScheduleResponse, error)
	CreateSchedule(schedule scheduleDto.CreateSchedule) (*scheduleDto.ScheduleResponse, error)
	UpdateSchedule(schedule scheduleDto.UpdateSchedule) (*scheduleDto.ScheduleResponse, error)
	DeleteSchedule(scheduleID int64) error
	UpdateScheduleSlot(scheduleSlotID int64, updatedSlot scheduleDto.CreateScheduleSlot) (*scheduleDto.ScheduleSlot, error)
	BindLessonToScheduleSlot(scheduleSlotID, lessonID int64) error
	DeleteLessonFromScheduleSlot(scheduleSlotID int64) error
}

func NewScheduleHandler(service ScheduleService, logger *zap.Logger) *ScheduleHandler {
	return &ScheduleHandler{
		service: service,
		logger:  logger,
	}
}

func (s *ScheduleHandler) GetAllSchedules(w http.ResponseWriter, r *http.Request) {
	res, err := s.service.GetAllSchedules()
	if err != nil {
		s.logger.Error("failed to get all schedules", zap.Error(err))
		http.Error(w, "failed to get all schedules", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func (s *ScheduleHandler) GetScheduleByID(w http.ResponseWriter, r *http.Request) {
	strScheduleID := chi.URLParam(r, "scheduleID")
	scheduleID, err := strconv.Atoi(strScheduleID)
	if err != nil {
		s.logger.Error("invalid param schedule id", zap.Error(err))
		http.Error(w, "invalid param schedule id", http.StatusBadRequest)
		return
	}

	res, err := s.service.GetScheduleByID(int64(scheduleID))
	if err != nil {
		s.logger.Error("failed to get schedule by id", zap.Int("scheduleID", scheduleID), zap.Error(err))
		http.Error(w, "failed to get schedule by id", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func (s *ScheduleHandler) GetSchedulesByTutorID(w http.ResponseWriter, r *http.Request) {
	strTutorID := chi.URLParam(r, "tutorID")
	tutorID, err := strconv.Atoi(strTutorID)
	if err != nil {
		s.logger.Error("invalid param tutor id", zap.Error(err))
		http.Error(w, "invalid param tutor id", http.StatusBadRequest)
		return
	}

	res, err := s.service.GetSchedulesByTutorID(int64(tutorID))
	if err != nil {
		s.logger.Error("failed to get tutor by id", zap.Int("tutorID", tutorID), zap.Error(err))
		http.Error(w, "failed to get tutor by id", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func (s *ScheduleHandler) CreateSchedule(w http.ResponseWriter, r *http.Request) {
	var schedule scheduleDto.CreateSchedule

	err := json.NewDecoder(r.Body).Decode(&schedule)
	if err != nil {
		s.logger.Error("invalid schedule body", zap.Error(err))
		http.Error(w, "invalid schedule body", http.StatusBadRequest)
		return
	}

	res, err := s.service.CreateSchedule(schedule)
	if err != nil {
		s.logger.Error("failed to create schedule", zap.Error(err))
		http.Error(w, "failed to create schedule", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func (s *ScheduleHandler) UpdateSchedule(w http.ResponseWriter, r *http.Request) {
	strScheduleID := chi.URLParam(r, "scheduleID")
	scheduleID, err := strconv.Atoi(strScheduleID)
	if err != nil {
		s.logger.Error("invalid param schedule id", zap.Error(err))
		http.Error(w, "invalid param schedule id", http.StatusBadRequest)
		return
	}

	var schedule scheduleDto.UpdateSchedule

	err = json.NewDecoder(r.Body).Decode(&schedule)
	if err != nil {
		s.logger.Error("invalid schedule body", zap.Error(err))
		http.Error(w, "invalid schedule body", http.StatusBadRequest)
		return
	}

	schedule.ID = int64(scheduleID)

	res, err := s.service.UpdateSchedule(schedule)
	if err != nil {
		s.logger.Error("failed to update schedule", zap.Error(err))
		http.Error(w, "failed to update schedule", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func (s *ScheduleHandler) DeleteSchedule(w http.ResponseWriter, r *http.Request) {
	strScheduleID := chi.URLParam(r, "scheduleID")
	scheduleID, err := strconv.Atoi(strScheduleID)
	if err != nil {
		s.logger.Error("invalid param schedule id", zap.Error(err))
		http.Error(w, "invalid param schedule id", http.StatusBadRequest)
		return
	}

	err = s.service.DeleteSchedule(int64(scheduleID))
	if err != nil {
		s.logger.Error("failed to delete schedule", zap.Int("scheduleID", scheduleID), zap.Error(err))
		http.Error(w, "failed to delete schedule", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (s *ScheduleHandler) UpdateScheduleSlot(w http.ResponseWriter, r *http.Request) {
	strScheduleSlotID := chi.URLParam(r, "scheduleSlotID")
	scheduleSlotID, err := strconv.Atoi(strScheduleSlotID)
	if err != nil {
		s.logger.Error("invalid param schedule slot id", zap.Error(err))
		http.Error(w, "invalid param schedule slot id", http.StatusBadRequest)
		return
	}

	var scheduleSlot scheduleDto.CreateScheduleSlot
	err = json.NewDecoder(r.Body).Decode(&scheduleSlot)
	if err != nil {
		s.logger.Error(
			"invalid update schedule slot dto",
			zap.Int("scheduleSlotID", scheduleSlotID),
			zap.Error(err),
		)
		http.Error(w, "invalid update schedule slot dto", http.StatusBadRequest)
		return
	}

	res, err := s.service.UpdateScheduleSlot(int64(scheduleSlotID), scheduleSlot)
	if err != nil {
		s.logger.Error(
			"failed to update schedule slot",
			zap.Int("scheduleSLotID", scheduleSlotID),
			zap.Error(err),
		)
		http.Error(w, "failed to update schedule slot", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func (s *ScheduleHandler) BindLessonToScheduleSlot(w http.ResponseWriter, r *http.Request) {
	strScheduleSlotID := chi.URLParam(r, "scheduleSlotID")
	scheduleSlotID, err := strconv.Atoi(strScheduleSlotID)
	if err != nil {
		s.logger.Error("invalid param schedule slot id", zap.Error(err))
		http.Error(w, "invalid param schedule slot id", http.StatusBadRequest)
		return
	}

	var lessonID int64
	err = json.NewDecoder(r.Body).Decode(&lessonID)
	if err != nil {
		s.logger.Error("invalid lesson id", zap.Error(err))
		http.Error(w, "invalid lesson id", http.StatusBadRequest)
		return
	}

	err = s.service.BindLessonToScheduleSlot(int64(scheduleSlotID), lessonID)
	if err != nil {
		s.logger.Error(
			"failed to bind lesson to schedule",
			zap.Int("scheduleSlotID", scheduleSlotID),
			zap.Int64("lessonID", lessonID),
			zap.Error(err),
		)
		http.Error(w, "invalid lesson id", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (s *ScheduleHandler) DeleteLessonFromScheduleSlot(w http.ResponseWriter, r *http.Request) {
	strScheduleSlotID := chi.URLParam(r, "scheduleSlotID")
	scheduleSlotID, err := strconv.Atoi(strScheduleSlotID)
	if err != nil {
		s.logger.Error("invalid param schedule slot id", zap.Error(err))
		http.Error(w, "invalid param schedule slot id", http.StatusBadRequest)
		return
	}

	err = s.service.DeleteLessonFromScheduleSlot(int64(scheduleSlotID))
	if err != nil {
		s.logger.Error(
			"failed to delete lesson from schedule slot",
			zap.Int("scheduleSlotID", scheduleSlotID),
			zap.Error(err),
		)
		http.Error(w, "failed to delete lesson from schedule slot", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
