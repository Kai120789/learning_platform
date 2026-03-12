package handler

import (
	"encoding/json"
	"go.uber.org/zap"
	"learning-platform/api-gateway/internal/config"
	"learning-platform/api-gateway/internal/dto"
	"learning-platform/api-gateway/internal/utils"
	"net/http"
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
		"access_token",
		res.AccessToken,
		time.Now().Add(time.Duration(a.cfg.AccessTokenLiveTime)*time.Minute),
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
		"access_token",
		res.AccessToken,
		time.Now().Add(time.Duration(a.cfg.AccessTokenLiveTime)*time.Minute),
	)

	http.SetCookie(w, cookie)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res)
}

func (a *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("access_token")
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

	cookie = utils.DeleteCookie("access_token")
	http.SetCookie(w, cookie)
	w.WriteHeader(http.StatusOK)
}
