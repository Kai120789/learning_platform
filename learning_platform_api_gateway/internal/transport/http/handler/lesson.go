package handler

import (
	"go.uber.org/zap"
	"net/http"
)

type LessonHandler struct {
	service LessonService
	logger  *zap.Logger
}

type LessonService interface{}

func NewLessonHandler(service LessonService, logger *zap.Logger) *LessonHandler {
	return &LessonHandler{
		service: service,
		logger:  logger,
	}
}

func (l *LessonHandler) GetOneLesson(w http.ResponseWriter, r *http.Request) {

}

func (l *LessonHandler) GetLessonsByUserId(w http.ResponseWriter, r *http.Request) {

}

func (l *LessonHandler) CreateLesson(w http.ResponseWriter, r *http.Request) {

}

func (l *LessonHandler) UpdateLesson(w http.ResponseWriter, r *http.Request) {

}

func (l *LessonHandler) UpdateLessonStatus(w http.ResponseWriter, r *http.Request) {

}

func (l *LessonHandler) GetLessonsByTutorId(w http.ResponseWriter, r *http.Request) {

}
