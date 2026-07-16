package handler

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
	"learning-platform/api-gateway/internal/dto/authDto"
	"learning-platform/api-gateway/internal/dto/enum"
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
	UpdateUserInfo(userID int64, userInfo userDto.UserInfoRequest) (*userDto.UserInfoResponse, error)
	UpdateUserSettings(userID int64, userSettings userDto.UserSettingsRequest) (*userDto.UserSettingsResponse, error)
	UpdateUserTheme(userID int64, theme enum.UserTheme) error
	UpdateUserAvatar(userID int64, avatar string) error
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
		u.logger.Error(
			"invalid param user id",
			zap.Int("userID", id),
			zap.Error(err),
		)
		http.Error(w, "invalid param user id", http.StatusBadRequest)
		return
	}

	user, err := u.service.GetUserById(int64(id))
	if err != nil {
		u.logger.Error(
			"failed get user by id",
			zap.Int("userID", id),
			zap.Error(err),
		)
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
		u.logger.Error(
			"failed get user data by id",
			zap.Int64("userID", userID),
			zap.Error(err),
		)
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
		u.logger.Error(
			"invalid create user dto",
			zap.String("email", createUserDto.Email),
			zap.Error(err),
		)
		http.Error(w, "invalid create user dto", http.StatusBadRequest)
		return
	}

	newUserId, err := u.service.CreateUser(createUserDto)
	if err != nil {
		u.logger.Error(
			"failed create user",
			zap.String("email", createUserDto.Email),
			zap.Error(err),
		)
		http.Error(w, "failed create user", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newUserId)
}

func (u *UserHandler) UpdateUserInfo(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value("user_id").(int64)
	if !ok {
		u.logger.Error(
			"user unauthorized",
			zap.Int64("userID", userID),
		)
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	var userInfo userDto.UserInfoRequest
	err := json.NewDecoder(r.Body).Decode(&userInfo)
	if err != nil {
		u.logger.Error(
			"invalid update user info dto",
			zap.Error(err),
		)
		http.Error(w, "invalid update user info dto", http.StatusBadRequest)
		return
	}

	res, err := u.service.UpdateUserInfo(userID, userInfo)
	if err != nil {
		u.logger.Error(
			"failed to update user info",
			zap.Int64("userID", userID),
			zap.Error(err),
		)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func (u *UserHandler) UpdateUserSettings(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value("user_id").(int64)
	if !ok {
		u.logger.Error(
			"user unauthorized",
			zap.Int64("userID", userID),
		)
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	var userSettings userDto.UserSettingsRequest
	err := json.NewDecoder(r.Body).Decode(&userSettings)
	if err != nil {
		u.logger.Error(
			"invalid update user settings dto",
			zap.Error(err),
		)
		http.Error(w, "invalid update user settings dto", http.StatusBadRequest)
		return
	}

	res, err := u.service.UpdateUserSettings(userID, userSettings)
	if err != nil {
		u.logger.Error(
			"failed to update user settings",
			zap.Int64("userID", userID),
			zap.Error(err),
		)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func (u *UserHandler) UpdateUserTheme(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value("user_id").(int64)
	if !ok {
		u.logger.Error(
			"user unauthorized",
			zap.Int64("userID", userID),
		)
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	theme := r.URL.Query().Get("theme")

	err := u.service.UpdateUserTheme(userID, enum.UserTheme(theme))
	if err != nil {
		u.logger.Error(
			"failed to update user theme",
			zap.Int64("userID", userID),
			zap.Error(err),
		)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (u *UserHandler) UpdateUserAvatar(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value("user_id").(int64)
	if !ok {
		u.logger.Error(
			"user unauthorized",
			zap.Int64("userID", userID),
		)
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	avatar := r.URL.Query().Get("avatar")

	err := u.service.UpdateUserAvatar(userID, avatar)
	if err != nil {
		u.logger.Error(
			"failed to update user avatar",
			zap.Int64("userID", userID),
			zap.Error(err),
		)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
