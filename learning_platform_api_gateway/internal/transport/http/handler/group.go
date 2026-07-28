package handler

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
	"learning-platform/api-gateway/internal/dto/groupDto"
	"net/http"
	"strconv"
)

type GroupHandler struct {
	service GroupService
	logger  *zap.Logger
}

type GroupService interface {
	CreateGroup(group groupDto.CreateGroupRequest, tutorID int64) (*groupDto.GroupFullResponse, error)
	UpdateGroup(groupID int64, newGroup groupDto.UpdateGroupRequest) (*groupDto.GroupResponse, error)
	RemoveGroup(groupID int64) error
	GetGroupByID(groupID int64) (*groupDto.GroupResponse, error)
	GetGroups() ([]groupDto.GroupResponse, error)
	AddUsersToGroup(groupID int64, userIDs []int64) ([]int64, error)
	RemoveUserFromGroup(userID int64, groupID int64) error
	GetGroupsByStudentID(userID int64) ([]groupDto.GroupFullResponse, error)
	GetGroupsByTutorID(tutorID int64) ([]groupDto.GroupFullResponse, error)
	GetGroupUsers(groupID int64) ([]int64, error)
}

func NewGroupHandler(service GroupService, logger *zap.Logger) *GroupHandler {
	return &GroupHandler{
		service: service,
		logger:  logger,
	}
}

func (g *GroupHandler) CreateGroup(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value("user_id").(int64)
	if !ok {
		g.logger.Error(
			"user unauthorized",
			zap.Int64("userID", userID),
		)
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	var group groupDto.CreateGroupRequest

	err := json.NewDecoder(r.Body).Decode(&group)
	if err != nil {
		g.logger.Error("invalid create group dto", zap.Error(err))
		http.Error(w, "invalid create group dto", http.StatusBadRequest)
		return
	}

	newGroup, err := g.service.CreateGroup(group, userID)
	if err != nil {
		g.logger.Error("failed to create group", zap.Error(err))
		http.Error(w, "failed to create group", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newGroup)
}

func (g *GroupHandler) UpdateGroup(w http.ResponseWriter, r *http.Request) {
	strGroupID := chi.URLParam(r, "groupID")

	groupID, err := strconv.Atoi(strGroupID)
	if err != nil {
		g.logger.Error("invalid param group id", zap.Error(err))
		http.Error(w, "invalid param group id", http.StatusBadRequest)
		return
	}

	var newGroup groupDto.UpdateGroupRequest

	err = json.NewDecoder(r.Body).Decode(&newGroup)
	if err != nil {
		g.logger.Error("invalid update group dto", zap.Error(err))
		http.Error(w, "invalid update group dto", http.StatusBadRequest)
		return
	}

	resGroup, err := g.service.UpdateGroup(int64(groupID), newGroup)
	if err != nil {
		g.logger.Error("failed to update group", zap.Error(err))
		http.Error(w, "failed to update group", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resGroup)
}

func (g *GroupHandler) RemoveGroup(w http.ResponseWriter, r *http.Request) {
	strGroupID := chi.URLParam(r, "groupID")

	groupID, err := strconv.Atoi(strGroupID)
	if err != nil {
		g.logger.Error("invalid param group id", zap.Error(err))
		http.Error(w, "invalid param group id", http.StatusBadRequest)
		return
	}

	err = g.service.RemoveGroup(int64(groupID))
	if err != nil {
		g.logger.Error("failed to delete group", zap.Error(err))
		http.Error(w, "failed to delete group", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (g *GroupHandler) GetGroupByID(w http.ResponseWriter, r *http.Request) {
	strGroupID := chi.URLParam(r, "groupID")

	groupID, err := strconv.Atoi(strGroupID)
	if err != nil {
		g.logger.Error("invalid param group id", zap.Error(err))
		http.Error(w, "invalid param group id", http.StatusBadRequest)
		return
	}

	group, err := g.service.GetGroupByID(int64(groupID))
	if err != nil {
		g.logger.Error("failed to get one group", zap.Error(err))
		http.Error(w, "failed to get one group", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(group)
}

func (g *GroupHandler) GetGroups(w http.ResponseWriter, r *http.Request) {
	groups, err := g.service.GetGroups()
	if err != nil {
		g.logger.Error("failed to get groups", zap.Error(err))
		http.Error(w, "failed to get groups", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(groups)
}

func (g *GroupHandler) AddUsersToGroup(w http.ResponseWriter, r *http.Request) {
	strGroupID := chi.URLParam(r, "groupID")

	groupID, err := strconv.Atoi(strGroupID)
	if err != nil {
		g.logger.Error("invalid param group id", zap.Error(err))
		http.Error(w, "invalid param group id", http.StatusBadRequest)
		return
	}

	var userIDs []int64

	err = json.NewDecoder(r.Body).Decode(&userIDs)
	if err != nil {
		g.logger.Error("invalid user ids", zap.Error(err))
		http.Error(w, "invalid user ids", http.StatusBadRequest)
		return
	}

	users, err := g.service.AddUsersToGroup(int64(groupID), userIDs)
	if err != nil {
		g.logger.Error("failed to add users to group", zap.Error(err))
		http.Error(w, "failed to add users to group", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

func (g *GroupHandler) RemoveUserFromGroup(w http.ResponseWriter, r *http.Request) {
	strGroupID := chi.URLParam(r, "groupID")

	groupID, err := strconv.Atoi(strGroupID)
	if err != nil {
		g.logger.Error("invalid param group id", zap.Error(err))
		http.Error(w, "invalid param group id", http.StatusBadRequest)
		return
	}

	strUserID := chi.URLParam(r, "userID")

	userID, err := strconv.Atoi(strUserID)
	if err != nil {
		g.logger.Error("invalid param user id", zap.Error(err))
		http.Error(w, "invalid param user id", http.StatusBadRequest)
		return
	}

	err = g.service.RemoveUserFromGroup(int64(userID), int64(groupID))
	if err != nil {
		g.logger.Error("failed to remove users from group", zap.Error(err))
		http.Error(w, "failed to remove users from group", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (g *GroupHandler) GetGroupsByStudentID(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value("user_id").(int64)
	if !ok {
		g.logger.Error(
			"user unauthorized",
			zap.Int64("userID", userID),
		)
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	groups, err := g.service.GetGroupsByStudentID(userID)
	if err != nil {
		g.logger.Error("failed to get user groups", zap.Error(err))
		http.Error(w, "failed to get user groups", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(groups)
}

func (g *GroupHandler) GetGroupsByTutorID(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value("user_id").(int64)
	if !ok {
		g.logger.Error(
			"user unauthorized",
			zap.Int64("userID", userID),
		)
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	groups, err := g.service.GetGroupsByTutorID(userID)
	if err != nil {
		g.logger.Error("failed to get tutor groups", zap.Error(err))
		http.Error(w, "failed to get tutor groups", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(groups)
}

func (g *GroupHandler) GetGroupUsers(w http.ResponseWriter, r *http.Request) {
	strGroupID := chi.URLParam(r, "groupID")

	groupID, err := strconv.Atoi(strGroupID)
	if err != nil {
		g.logger.Error("invalid param user id", zap.Error(err))
		http.Error(w, "invalid param user id", http.StatusBadRequest)
		return
	}

	users, err := g.service.GetGroupUsers(int64(groupID))
	if err != nil {
		g.logger.Error("failed to get group users", zap.Error(err))
		http.Error(w, "failed to get group users", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}
