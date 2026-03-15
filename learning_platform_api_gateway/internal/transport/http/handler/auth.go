package handler

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
	"learning-platform/api-gateway/internal/config"
	"learning-platform/api-gateway/internal/dto"
	"learning-platform/api-gateway/internal/utils"
	"net/http"
	"strconv"
	"time"
)

type AuthHandler struct {
	service AuthService
	logger  *zap.Logger
	cfg     *config.Config
}

type AuthService interface {
	Login(loginReq dto.LoginRequest) (*dto.LoginResponse, error)
	Register(registerReq dto.RegisterRequest) (*dto.RegisterResponse, error)
	Logout(accessToken string) error
	LogoutAll(userId int64) error
}

func NewAuthHandler(
	service AuthService,
	logger *zap.Logger,
	cfg *config.Config,
) *AuthHandler {
	return &AuthHandler{
		service: service,
		logger:  logger,
		cfg:     cfg,
	}
}

func (a *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var loginReq dto.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&loginReq); err != nil {
		http.Error(w, "invalid input", http.StatusBadRequest)
		return
	}

	res, err := a.service.Login(loginReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	cookie := utils.CreateCookie(
		"session_id",
		res.SessionId,
		time.Now().Add(time.Duration(a.cfg.RefreshTokenLiveTime)*time.Hour*24),
	)
	http.SetCookie(w, cookie)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func (a *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var registerReq dto.RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&registerReq); err != nil {
		http.Error(w, "invalid input", http.StatusBadRequest)
		return
	}

	res, err := a.service.Register(registerReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	cookie := utils.CreateCookie(
		"session_id",
		res.SessionId,
		time.Now().Add(time.Duration(a.cfg.RefreshTokenLiveTime)*time.Hour*24),
	)
	http.SetCookie(w, cookie)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res)
}

func (a *AuthHandler) RefreshTokens(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("refresh")
}

func (a *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session_id")
	if err != nil {
		http.Error(w, "cookie not found", http.StatusNotFound)
		return
	}

	if cookie.Value == "" {
		http.Error(w, "incorrect cookie value", http.StatusBadRequest)
		return
	}

	err = a.service.Logout(cookie.Value)
	if err != nil {
		http.Error(w, "failed logout user", http.StatusInternalServerError)
		return
	}

	cookie = utils.DeleteCookie("session_id")
	http.SetCookie(w, cookie)
	w.WriteHeader(http.StatusOK)
}

func (a *AuthHandler) LogoutAll(w http.ResponseWriter, r *http.Request) {
	strUserId := chi.URLParam(r, "userId")
	if strUserId == "" {
		http.Error(w, "invalid param user id", http.StatusBadRequest)
		return
	}

	userId, err := strconv.Atoi(strUserId)
	if err != nil {
		http.Error(w, "error convert string param to int", http.StatusInternalServerError)
		return
	}

	err = a.service.LogoutAll(int64(userId))
	if err != nil {
		http.Error(w, "failed logout all", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
