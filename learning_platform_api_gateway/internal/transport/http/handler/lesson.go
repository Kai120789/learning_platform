package handler

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
	"learning-platform/api-gateway/internal/dto/enum"
	"learning-platform/api-gateway/internal/dto/lessonDto"
	"net/http"
	"strconv"
)

type LessonHandler struct {
	service LessonService
	logger  *zap.Logger
}

type LessonService interface {
	GetOneLesson(lessonID int64) (*lessonDto.LessonResponse, error)
	GetLessonsByUserId(userID int64) ([]lessonDto.LessonResponse, error)
	CreateLesson(lesson lessonDto.CreateLesson) (*lessonDto.LessonResponse, error)
	UpdateLesson(lesson lessonDto.UpdateLesson) (*lessonDto.LessonResponse, error)
	UpdateLessonStatus(lessonID int64, status enum.LessonStatus) error
	GetLessonsByTutorId(tutorID int64) ([]lessonDto.LessonResponse, error)
}

func NewLessonHandler(service LessonService, logger *zap.Logger) *LessonHandler {
	return &LessonHandler{
		service: service,
		logger:  logger,
	}
}

func (l *LessonHandler) GetOneLesson(w http.ResponseWriter, r *http.Request) {
	strLessonID := chi.URLParam(r, "lessonId")
	lessonID, err := strconv.Atoi(strLessonID)
	if err != nil {
		l.logger.Error("invalid param lesson id", zap.Error(err))
		http.Error(w, "invalid param lesson id", http.StatusBadRequest)
		return
	}

	res, err := l.service.GetOneLesson(int64(lessonID))
	if err != nil {
		l.logger.Error("failed to get one lesson", zap.Int("lessonID", lessonID), zap.Error(err))
		http.Error(w, "failed to get one lesson", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func (l *LessonHandler) GetLessonsByUserId(w http.ResponseWriter, r *http.Request) {
	strUserID := chi.URLParam(r, "userID")
	userID, err := strconv.Atoi(strUserID)
	if err != nil {
		l.logger.Error("invalid param user id", zap.Error(err))
		http.Error(w, "invalid param user id", http.StatusBadRequest)
		return
	}

	res, err := l.service.GetLessonsByUserId(int64(userID))
	if err != nil {
		l.logger.Error("failed to get lessons by user id", zap.Int("userID", userID), zap.Error(err))
		http.Error(w, "failed to get lessons by user id", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func (l *LessonHandler) CreateLesson(w http.ResponseWriter, r *http.Request) {
	var lesson lessonDto.CreateLesson
	err := json.NewDecoder(r.Body).Decode(&lesson)
	if err != nil {
		l.logger.Error("invalid lesson dto", zap.Error(err))
		http.Error(w, "invalid lesson dto", http.StatusBadRequest)
		return
	}

	res, err := l.service.CreateLesson(lesson)
	if err != nil {
		l.logger.Error("failed to create lesson", zap.Error(err))
		http.Error(w, "failed to create lesson", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res)
}

func (l *LessonHandler) UpdateLesson(w http.ResponseWriter, r *http.Request) {
	var lesson lessonDto.UpdateLesson
	err := json.NewDecoder(r.Body).Decode(&lesson)
	if err != nil {
		l.logger.Error("invalid lesson dto", zap.Error(err))
		http.Error(w, "invalid lesson dto", http.StatusBadRequest)
		return
	}

	res, err := l.service.UpdateLesson(lesson)
	if err != nil {
		l.logger.Error("failed to update lesson", zap.Error(err))
		http.Error(w, "failed to update lesson", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func (l *LessonHandler) UpdateLessonStatus(w http.ResponseWriter, r *http.Request) {
	strLessonID := chi.URLParam(r, "lessonId")
	lessonID, err := strconv.Atoi(strLessonID)
	if err != nil {
		l.logger.Error("invalid param lesson id", zap.Error(err))
		http.Error(w, "invalid param lesson id", http.StatusBadRequest)
		return
	}

	var newStatus enum.LessonStatus
	err = json.NewDecoder(r.Body).Decode(&newStatus)
	if err != nil {
		l.logger.Error("invalid lesson status value", zap.Error(err))
		http.Error(w, "invalid lesson status value", http.StatusBadRequest)
		return
	}

	err = l.service.UpdateLessonStatus(int64(lessonID), newStatus)
	if err != nil {
		l.logger.Error("failed to update lesson status", zap.Int("lessonID", lessonID), zap.Error(err))
		http.Error(w, "failed to update lesson status", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (l *LessonHandler) GetLessonsByTutorId(w http.ResponseWriter, r *http.Request) {
	strTutorID := chi.URLParam(r, "tutorID")
	tutorID, err := strconv.Atoi(strTutorID)
	if err != nil {
		l.logger.Error("invalid param tutor id", zap.Error(err))
		http.Error(w, "invalid param tutor id", http.StatusBadRequest)
		return
	}

	res, err := l.service.GetLessonsByTutorId(int64(tutorID))
	if err != nil {
		l.logger.Error("failed to get lessons by tutor id", zap.Int("tutorID", tutorID), zap.Error(err))
		http.Error(w, "failed to get lessons by tutor id", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}
