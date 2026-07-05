package handler

import (
	"go.uber.org/zap"
	"net/http"
)

type SubjectHandler struct {
	service SubjectService
	logger  *zap.Logger
}

type SubjectService interface{}

func NewSubjectHandler(service SubjectService, logger *zap.Logger) *SubjectHandler {
	return &SubjectHandler{
		service: service,
		logger:  logger,
	}
}

func (s *SubjectHandler) GetOneSubject(w http.ResponseWriter, r *http.Request) {

}

func (s *SubjectHandler) GetAllSubjects(w http.ResponseWriter, r *http.Request) {

}

func (s *SubjectHandler) GetUserSubjects(w http.ResponseWriter, r *http.Request) {

}

func (s *SubjectHandler) SetUserSubjects(w http.ResponseWriter, r *http.Request) {

}

func (s *SubjectHandler) UpdateUserSubjects(w http.ResponseWriter, r *http.Request) {

}
