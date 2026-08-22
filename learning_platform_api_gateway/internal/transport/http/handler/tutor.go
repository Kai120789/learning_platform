package handler

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
	"learning-platform/api-gateway/internal/dto/tutorDto"
	"net/http"
	"strconv"
)

type TutorHandler struct {
	service TutorService
	logger  *zap.Logger
}

type TutorService interface {
	AddStudent(tutorID, studentID int64) error
	DeleteOneStudent(tutorID, studentID int64) error
	DeleteStudents(tutorID int64, studentIDs []int64) error
	GetTutorStudents(getTutorStudents tutorDto.GetTutorStudents) ([]tutorDto.TutorStudentWithInfo, int64, error)
	GetOneTutor(tutorID, userID int64) (*tutorDto.OneTutorWithInfo, error)
	GetTutors(getTutors tutorDto.GetTutors) ([]tutorDto.TutorShortInfoWithSubjects, int64, error)
	GetTutorReviews(getTutorReviews tutorDto.GetTutorReviews) ([]tutorDto.TutorReviewWithInfo, int64, error)
	AddTutorReview(newReview tutorDto.NewTutorReview) (*tutorDto.TutorReviewWithInfo, error)
	UpdateTutorReview(updReview tutorDto.UpdateTutorReview) (*tutorDto.TutorReviewWithInfo, error)
	DeleteTutorReview(reviewID, authorID int64) error
	GetTutorOffers(tutorID int64) ([]tutorDto.TutorOfferWithInfo, error)
	AddTutorOffer(newOffer tutorDto.NewTutorOffer) (*tutorDto.TutorOfferWithInfo, error)
	UpdateTutorOffer(updOffer tutorDto.UpdateTutorOffer) (*tutorDto.TutorOfferWithInfo, error)
	DeleteOneTutorOffer(offerID, tutorID int64) error
	DeleteTutorOffers(offerIDs []int64, tutorID int64) error
	AddTutorSubjects(tutorID int64, subjectIDs []int64) ([]int64, error)
	UpdateTutorSubjects(tutorID int64, subjectIDs []int64) ([]int64, error)
}

func NewTutorHandler(service TutorService, logger *zap.Logger) *TutorHandler {
	return &TutorHandler{
		service: service,
		logger:  logger,
	}
}

func (t *TutorHandler) AddStudent(w http.ResponseWriter, r *http.Request) {
	strStudentID := chi.URLParam(r, "studentID")
	studentID, err := strconv.Atoi(strStudentID)
	if err != nil {
		t.logger.Error("invalid param student id", zap.Error(err))
		http.Error(w, "invalid param student id", http.StatusBadRequest)
		return
	}

	userID, ok := r.Context().Value("user_id").(int64)
	if !ok {
		t.logger.Error(
			"user unauthorized",
			zap.Int64("userID", userID),
		)
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	err = t.service.AddStudent(userID, int64(studentID))
	if err != nil {
		t.logger.Error(
			"failed to add tutor student",
			zap.Int64("tutorID", userID),
			zap.Int("studentID", studentID),
			zap.Error(err),
		)
		http.Error(w, "failed to add tutor student", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (t *TutorHandler) DeleteOneStudent(w http.ResponseWriter, r *http.Request) {
	strStudentID := chi.URLParam(r, "studentID")
	studentID, err := strconv.Atoi(strStudentID)
	if err != nil {
		t.logger.Error("invalid param student id", zap.Error(err))
		http.Error(w, "invalid param student id", http.StatusBadRequest)
		return
	}

	userID, ok := r.Context().Value("user_id").(int64)
	if !ok {
		t.logger.Error(
			"user unauthorized",
			zap.Int64("userID", userID),
		)
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	err = t.service.DeleteOneStudent(userID, int64(studentID))
	if err != nil {
		t.logger.Error(
			"failed to delete one tutor student",
			zap.Int64("tutorID", userID),
			zap.Int("studentID", studentID),
			zap.Error(err),
		)
		http.Error(w, "failed to delete one tutor student", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (t *TutorHandler) DeleteStudents(w http.ResponseWriter, r *http.Request) {
	var studentIDs []int64
	err := json.NewDecoder(r.Body).Decode(&studentIDs)
	if err != nil {
		t.logger.Error(
			"invalid student IDs",
			zap.Error(err),
		)
		http.Error(w, "invalid student IDs", http.StatusBadRequest)
		return
	}

	userID, ok := r.Context().Value("user_id").(int64)
	if !ok {
		t.logger.Error(
			"user unauthorized",
			zap.Int64("userID", userID),
		)
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	err = t.service.DeleteStudents(userID, studentIDs)
	if err != nil {
		t.logger.Error(
			"failed to delete tutor students",
			zap.Int64("tutorID", userID),
			zap.Int64s("studentIDs", studentIDs),
			zap.Error(err),
		)
		http.Error(w, "failed to delete tutor students", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (t *TutorHandler) GetTutorStudents(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value("user_id").(int64)
	if !ok {
		t.logger.Error(
			"user unauthorized",
			zap.Int64("userID", userID),
		)
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	pageStr := r.URL.Query().Get("page")
	limitStr := r.URL.Query().Get("limit")
	interactedWithinDaysStr := r.URL.Query().Get("interacted_within_days")

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		t.logger.Error(
			"invalid param page",
			zap.Error(err),
		)
		http.Error(w, "invalid param page", http.StatusBadRequest)
		return
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		t.logger.Error(
			"invalid param limit",
			zap.Error(err),
		)
		http.Error(w, "invalid param limit", http.StatusBadRequest)
		return
	}

	var interactedWithinDaysOpt *int64
	if interactedWithinDaysStr != "" {
		interactedWithinDays, err := strconv.Atoi(interactedWithinDaysStr)
		if err != nil {
			t.logger.Error(
				"invalid param interacted_within_days",
				zap.Error(err),
			)
			http.Error(w, "invalid param interacted_within_days", http.StatusBadRequest)
			return
		}

		IWD64 := int64(interactedWithinDays)
		interactedWithinDaysOpt = &IWD64
	}

	students, count, err := t.service.GetTutorStudents(tutorDto.GetTutorStudents{
		TutorID:              userID,
		Page:                 int64(page),
		Limit:                int64(limit),
		InteractedWithinDays: interactedWithinDaysOpt,
	})
	if err != nil {
		t.logger.Error(
			"failed to get tutor students",
			zap.Int64("tutorID", userID),
			zap.Error(err),
		)
		http.Error(w, "failed to get tutor students", http.StatusInternalServerError)
		return
	}

	type responseDTO struct {
		Students []tutorDto.TutorStudentWithInfo `json:"students"`
		Count    int64                           `json:"count"`
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(responseDTO{
		Students: students,
		Count:    count,
	})
}

func (t *TutorHandler) GetOneTutor(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value("user_id").(int64)
	if !ok {
		t.logger.Error(
			"user unauthorized",
			zap.Int64("userID", userID),
		)
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	strTutorID := chi.URLParam(r, "tutorID")
	tutorID, err := strconv.Atoi(strTutorID)
	if err != nil {
		t.logger.Error("invalid param tutor id", zap.Error(err))
		http.Error(w, "invalid param tutor id", http.StatusBadRequest)
		return
	}

	res, err := t.service.GetOneTutor(int64(tutorID), userID)
	if err != nil {
		t.logger.Error(
			"failed to get one tutor",
			zap.Int64("userID", userID),
			zap.Int("tutorID", tutorID),
			zap.Error(err),
		)
		http.Error(w, "failed to get one tutor", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func (t *TutorHandler) GetTutors(w http.ResponseWriter, r *http.Request) {
	pageStr := r.URL.Query().Get("page")
	limitStr := r.URL.Query().Get("limit")
	subjectIDStr := r.URL.Query().Get("subject_id")

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		t.logger.Error(
			"invalid param page",
			zap.Error(err),
		)
		http.Error(w, "invalid param page", http.StatusBadRequest)
		return
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		t.logger.Error(
			"invalid param limit",
			zap.Error(err),
		)
		http.Error(w, "invalid param limit", http.StatusBadRequest)
		return
	}

	var subjectIDOpt *int64
	if subjectIDStr != "" {
		subjectID, err := strconv.Atoi(subjectIDStr)
		if err != nil {
			t.logger.Error(
				"invalid param subject_id",
				zap.Error(err),
			)
			http.Error(w, "invalid param subject_id", http.StatusBadRequest)
			return
		}

		SID64 := int64(subjectID)
		subjectIDOpt = &SID64
	}

	tutors, count, err := t.service.GetTutors(tutorDto.GetTutors{
		SubjectID: subjectIDOpt,
		Page:      int64(page),
		Limit:     int64(limit),
	})
	if err != nil {
		t.logger.Error(
			"failed to get tutors",
			zap.Error(err),
		)
		http.Error(w, "failed to get tutors", http.StatusInternalServerError)
		return
	}

	type responseDTO struct {
		Tutors []tutorDto.TutorShortInfoWithSubjects `json:"tutors"`
		Count  int64                                 `json:"count"`
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(responseDTO{
		Tutors: tutors,
		Count:  count,
	})
}

func (t *TutorHandler) GetTutorReviews(w http.ResponseWriter, r *http.Request) {
	strTutorID := chi.URLParam(r, "tutorID")
	tutorID, err := strconv.Atoi(strTutorID)
	if err != nil {
		t.logger.Error("invalid param tutor id", zap.Error(err))
		http.Error(w, "invalid param tutor id", http.StatusBadRequest)
		return
	}

	pageStr := r.URL.Query().Get("page")
	limitStr := r.URL.Query().Get("limit")

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		t.logger.Error(
			"invalid param page",
			zap.Error(err),
		)
		http.Error(w, "invalid param page", http.StatusBadRequest)
		return
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		t.logger.Error(
			"invalid param limit",
			zap.Error(err),
		)
		http.Error(w, "invalid param limit", http.StatusBadRequest)
		return
	}

	reviews, count, err := t.service.GetTutorReviews(tutorDto.GetTutorReviews{
		TutorID: int64(tutorID),
		Page:    int64(page),
		Limit:   int64(limit),
	})
	if err != nil {
		t.logger.Error(
			"failed to get tutor reviews",
			zap.Error(err),
		)
		http.Error(w, "failed to get tutor reviews", http.StatusInternalServerError)
		return
	}

	type responseDTO struct {
		Reviews []tutorDto.TutorReviewWithInfo `json:"reviews"`
		Count   int64                          `json:"count"`
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(responseDTO{
		Reviews: reviews,
		Count:   count,
	})
}

func (t *TutorHandler) AddTutorReview(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value("user_id").(int64)
	if !ok {
		t.logger.Error(
			"user unauthorized",
			zap.Int64("userID", userID),
		)
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	var newTutorReview tutorDto.NewTutorReview
	err := json.NewDecoder(r.Body).Decode(&newTutorReview)
	if err != nil {
		t.logger.Error(
			"invalid add tutor review dto",
			zap.Error(err),
		)
		http.Error(w, "invalid add tutor review dto", http.StatusBadRequest)
		return
	}

	newTutorReview.AuthorID = userID

	res, err := t.service.AddTutorReview(newTutorReview)
	if err != nil {
		t.logger.Error(
			"failed to add tutor review",
			zap.Error(err),
		)
		http.Error(w, "failed to add tutor review", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res)
}

func (t *TutorHandler) UpdateTutorReview(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value("user_id").(int64)
	if !ok {
		t.logger.Error(
			"user unauthorized",
			zap.Int64("userID", userID),
		)
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	var updTutorReview tutorDto.UpdateTutorReview
	err := json.NewDecoder(r.Body).Decode(&updTutorReview)
	if err != nil {
		t.logger.Error(
			"invalid update tutor review dto",
			zap.Error(err),
		)
		http.Error(w, "invalid update tutor review dto", http.StatusBadRequest)
		return
	}

	updTutorReview.AuthorID = userID

	res, err := t.service.UpdateTutorReview(updTutorReview)
	if err != nil {
		t.logger.Error(
			"failed to update tutor review",
			zap.Error(err),
		)
		http.Error(w, "failed to update tutor review", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res)
}

func (t *TutorHandler) DeleteTutorReview(w http.ResponseWriter, r *http.Request) {
	strReviewID := chi.URLParam(r, "reviewID")
	reviewID, err := strconv.Atoi(strReviewID)
	if err != nil {
		t.logger.Error("invalid param review id", zap.Error(err))
		http.Error(w, "invalid param review id", http.StatusBadRequest)
		return
	}

	userID, ok := r.Context().Value("user_id").(int64)
	if !ok {
		t.logger.Error(
			"user unauthorized",
			zap.Int64("userID", userID),
		)
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	err = t.service.DeleteTutorReview(int64(reviewID), userID)
	if err != nil {
		t.logger.Error(
			"failed to delete tutor review",
			zap.Int64("userID", userID),
			zap.Int("reviewID", reviewID),
			zap.Error(err),
		)
		http.Error(w, "failed to delete tutor review", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (t *TutorHandler) GetTutorOffers(w http.ResponseWriter, r *http.Request) {
	strTutorID := chi.URLParam(r, "tutorID")
	tutorID, err := strconv.Atoi(strTutorID)
	if err != nil {
		t.logger.Error("invalid param tutor id", zap.Error(err))
		http.Error(w, "invalid param tutor id", http.StatusBadRequest)
		return
	}

	offers, err := t.service.GetTutorOffers(int64(tutorID))
	if err != nil {
		t.logger.Error(
			"failed to get tutor offers",
			zap.Error(err),
		)
		http.Error(w, "failed to get tutor offers", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(offers)
}

func (t *TutorHandler) AddTutorOffer(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value("user_id").(int64)
	if !ok {
		t.logger.Error(
			"user unauthorized",
			zap.Int64("userID", userID),
		)
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	var newTutorOffer tutorDto.NewTutorOffer
	err := json.NewDecoder(r.Body).Decode(&newTutorOffer)
	if err != nil {
		t.logger.Error(
			"invalid add tutor offer dto",
			zap.Error(err),
		)
		http.Error(w, "invalid add tutor offer dto", http.StatusBadRequest)
		return
	}

	newTutorOffer.TutorID = userID

	res, err := t.service.AddTutorOffer(newTutorOffer)
	if err != nil {
		t.logger.Error(
			"failed to add tutor offer",
			zap.Error(err),
		)
		http.Error(w, "failed to add tutor offer", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res)
}

func (t *TutorHandler) UpdateTutorOffer(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value("user_id").(int64)
	if !ok {
		t.logger.Error(
			"user unauthorized",
			zap.Int64("userID", userID),
		)
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	var updTutorOffer tutorDto.UpdateTutorOffer
	err := json.NewDecoder(r.Body).Decode(&updTutorOffer)
	if err != nil {
		t.logger.Error(
			"invalid update tutor offer dto",
			zap.Error(err),
		)
		http.Error(w, "invalid update tutor offer dto", http.StatusBadRequest)
		return
	}

	updTutorOffer.TutorID = userID

	res, err := t.service.UpdateTutorOffer(updTutorOffer)
	if err != nil {
		t.logger.Error(
			"failed to update tutor offer",
			zap.Error(err),
		)
		http.Error(w, "failed to update tutor offer", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res)
}

func (t *TutorHandler) DeleteOneTutorOffer(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value("user_id").(int64)
	if !ok {
		t.logger.Error(
			"user unauthorized",
			zap.Int64("userID", userID),
		)
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	strOfferID := chi.URLParam(r, "offerID")
	offerID, err := strconv.Atoi(strOfferID)
	if err != nil {
		t.logger.Error("invalid param offer id", zap.Error(err))
		http.Error(w, "invalid param offer id", http.StatusBadRequest)
		return
	}

	err = t.service.DeleteOneTutorOffer(int64(offerID), userID)
	if err != nil {
		t.logger.Error(
			"failed to delete one tutor offer",
			zap.Int64("tutorID", userID),
			zap.Int("offerID", offerID),
			zap.Error(err),
		)
		http.Error(w, "failed to delete one tutor offer", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (t *TutorHandler) DeleteTutorOffers(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value("user_id").(int64)
	if !ok {
		t.logger.Error(
			"user unauthorized",
			zap.Int64("userID", userID),
		)
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	var offerIDs []int64
	err := json.NewDecoder(r.Body).Decode(&offerIDs)
	if err != nil {
		t.logger.Error(
			"invalid offer IDs",
			zap.Error(err),
		)
		http.Error(w, "invalid offer IDs", http.StatusBadRequest)
		return
	}

	err = t.service.DeleteTutorOffers(offerIDs, userID)
	if err != nil {
		t.logger.Error(
			"failed to delete tutor offers",
			zap.Int64("tutorID", userID),
			zap.Int64s("offerIDs", offerIDs),
			zap.Error(err),
		)
		http.Error(w, "failed to delete tutor offers", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (t *TutorHandler) AddTutorSubjects(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value("user_id").(int64)
	if !ok {
		t.logger.Error(
			"user unauthorized",
			zap.Int64("userID", userID),
		)
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	var subjectIDs []int64
	err := json.NewDecoder(r.Body).Decode(&subjectIDs)
	if err != nil {
		t.logger.Error(
			"invalid subject IDs",
			zap.Error(err),
		)
		http.Error(w, "invalid subject IDs", http.StatusBadRequest)
		return
	}

	res, err := t.service.AddTutorSubjects(userID, subjectIDs)
	if err != nil {
		t.logger.Error(
			"failed to add tutor subjects",
			zap.Int64("tutorID", userID),
			zap.Int64s("subjectIDs", subjectIDs),
			zap.Error(err),
		)
		http.Error(w, "failed to add tutor subjects", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res)
}

func (t *TutorHandler) UpdateTutorSubjects(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value("user_id").(int64)
	if !ok {
		t.logger.Error(
			"user unauthorized",
			zap.Int64("userID", userID),
		)
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	var subjectIDs []int64
	err := json.NewDecoder(r.Body).Decode(&subjectIDs)
	if err != nil {
		t.logger.Error(
			"invalid subject IDs",
			zap.Error(err),
		)
		http.Error(w, "invalid subject IDs", http.StatusBadRequest)
		return
	}

	res, err := t.service.UpdateTutorSubjects(userID, subjectIDs)
	if err != nil {
		t.logger.Error(
			"failed to update tutor subjects",
			zap.Int64("tutorID", userID),
			zap.Int64s("subjectIDs", subjectIDs),
			zap.Error(err),
		)
		http.Error(w, "failed to update tutor subjects", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}
