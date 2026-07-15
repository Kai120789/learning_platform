package handler

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
	"learning-platform/api-gateway/internal/dto/authDto"
	"learning-platform/api-gateway/internal/dto/userDto"
	"net/http"
	"strconv"
)

type UserHandler struct {
	service UserService
	logger  *zap.Logger
}

type UserService interface {
	GetUserById(id int64) (*userDto.GetUser, error)
	GetUserData(id int64) (*userDto.UserData, error)
	CreateUser(newUser authDto.RegisterRequest) (*int64, error)
}

func NewUserHandler(service UserService, logger *zap.Logger) *UserHandler {
	return &UserHandler{
		service: service,
		logger:  logger,
	}
}

func (u *UserHandler) GetUserById(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "userId")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		u.logger.Error("invalid param user id")
		http.Error(w, "invalid param user id", http.StatusBadRequest)
		return
	}

	user, err := u.service.GetUserById(int64(id))
	if err != nil {
		u.logger.Error("failed get user by id")
		http.Error(w, "failed get user by id", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func (u *UserHandler) GetUserData(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value("user_id").(int64)
	if !ok {
		u.logger.Error(
			"user unauthorized",
			zap.Int64("userID", userID),
		)
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	userData, err := u.service.GetUserData(userID)
	if err != nil {
		u.logger.Error("failed get user data by id")
		http.Error(w, "failed get user data by id", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(userData)
}

func (u *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var createUserDto authDto.RegisterRequest
	err := json.NewDecoder(r.Body).Decode(&createUserDto)
	if err != nil {
		u.logger.Error("invalid create user dto")
		http.Error(w, "invalid create user dto", http.StatusBadRequest)
		return
	}

	newUserId, err := u.service.CreateUser(createUserDto)
	if err != nil {
		u.logger.Error("failed create user")
		http.Error(w, "failed create user", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newUserId)
}
