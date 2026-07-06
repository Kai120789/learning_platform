package handler

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
	"learning-platform/api-gateway/internal/dto/subjectDto"
	"net/http"
	"strconv"
)

type SubjectHandler struct {
	service SubjectService
	logger  *zap.Logger
}

type SubjectService interface {
	GetOneSubject(subjectID int64) (*subjectDto.Subject, error)
	GetAllSubjects() ([]subjectDto.Subject, error)
	GetUserSubjects(userID int64) ([]subjectDto.Subject, error)
	SetUserSubjects(userID int64, subjectIDs []int64) error
	UpdateUserSubjects(userID int64, subjectIDs, deletedSubjectIDs []int64) error
}

func NewSubjectHandler(service SubjectService, logger *zap.Logger) *SubjectHandler {
	return &SubjectHandler{
		service: service,
		logger:  logger,
	}
}

func (s *SubjectHandler) GetOneSubject(w http.ResponseWriter, r *http.Request) {
	strSubjectID := chi.URLParam(r, "subjectID")
	subjectID, err := strconv.Atoi(strSubjectID)
	if err != nil {
		s.logger.Error("invalid param subject id", zap.Error(err))
		http.Error(w, "invalid param subject id", http.StatusBadRequest)
		return
	}

	res, err := s.service.GetOneSubject(int64(subjectID))
	if err != nil {
		s.logger.Error(
			"failed to get one subject",
			zap.Int("subjectID", subjectID),
			zap.Error(err),
		)
		http.Error(w, "failed to get one subject", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func (s *SubjectHandler) GetAllSubjects(w http.ResponseWriter, r *http.Request) {
	res, err := s.service.GetAllSubjects()
	if err != nil {
		s.logger.Error("failed to get all subjects", zap.Error(err))
		http.Error(w, "failed to get all subjects", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func (s *SubjectHandler) GetUserSubjects(w http.ResponseWriter, r *http.Request) {
	strUserID := chi.URLParam(r, "userID")
	userID, err := strconv.Atoi(strUserID)
	if err != nil {
		s.logger.Error("invalid param user id", zap.Error(err))
		http.Error(w, "invalid param user id", http.StatusBadRequest)
		return
	}

	res, err := s.service.GetUserSubjects(int64(userID))
	if err != nil {
		s.logger.Error(
			"failed to get user subjects",
			zap.Int("userID", userID),
			zap.Error(err),
		)
		http.Error(w, "failed to get user subjects", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func (s *SubjectHandler) SetUserSubjects(w http.ResponseWriter, r *http.Request) {
	strUserID := chi.URLParam(r, "userID")
	userID, err := strconv.Atoi(strUserID)
	if err != nil {
		s.logger.Error("invalid param user id", zap.Error(err))
		http.Error(w, "invalid param user id", http.StatusBadRequest)
		return
	}

	var subjectIDs []int64
	err = json.NewDecoder(r.Body).Decode(&subjectIDs)
	if err != nil {
		s.logger.Error("invalid subject ids", zap.Error(err))
		http.Error(w, "invalid subject ids", http.StatusBadRequest)
		return
	}

	err = s.service.SetUserSubjects(int64(userID), subjectIDs)
	if err != nil {
		s.logger.Error(
			"failed set user subjects",
			zap.Int("userID", userID),
			zap.Error(err),
		)
		http.Error(w, "failed set user subjects", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (s *SubjectHandler) UpdateUserSubjects(w http.ResponseWriter, r *http.Request) {
	strUserID := chi.URLParam(r, "userID")
	userID, err := strconv.Atoi(strUserID)
	if err != nil {
		s.logger.Error("invalid param user id", zap.Error(err))
		http.Error(w, "invalid param user id", http.StatusBadRequest)
		return
	}

	var updateUserSubjects subjectDto.UpdateUserSubjects

	err = json.NewDecoder(r.Body).Decode(&updateUserSubjects)
	if err != nil {
		s.logger.Error("invalid subject ids", zap.Error(err))
		http.Error(w, "invalid subject ids", http.StatusBadRequest)
		return
	}

	err = s.service.UpdateUserSubjects(
		int64(userID),
		updateUserSubjects.SubjectIDs,
		updateUserSubjects.DeletedSubjectIDs,
	)
	if err != nil {
		s.logger.Error(
			"failed update user subjects",
			zap.Int("userID", userID),
			zap.Int64s("subjectIDs", updateUserSubjects.SubjectIDs),
			zap.Int64s("deletedSubjectIDs", updateUserSubjects.DeletedSubjectIDs),
			zap.Error(err),
		)
		http.Error(w, "failed set user subjects", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
