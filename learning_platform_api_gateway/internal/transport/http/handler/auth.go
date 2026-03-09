package handler

import (
	"encoding/json"
	"go.uber.org/zap"
	"learning-platform/api-gateway/internal/dto"
	"net/http"
)

type AuthHandler struct {
	service AuthService
	logger  *zap.Logger
}

type AuthService interface {
	Login(loginReq dto.LoginRequest) (*dto.LoginResponse, error)
	Register(registerReq dto.RegisterRequest) (*dto.RegisterResponse, error)
	RefreshTokens(accessToken string) (*string, error)
	Logout(accessToken string) error
}

func NewAuthHandler(service AuthService, logger *zap.Logger) *AuthHandler {
	return &AuthHandler{
		service: service,
		logger:  logger,
	}
}

func (a *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var loginReq dto.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&loginReq); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	res, err := a.service.Login(loginReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func (a *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var registerReq dto.RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&registerReq); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	res, err := a.service.Register(registerReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res)
}

func (a *AuthHandler) RefreshTokens(w http.ResponseWriter, r *http.Request) {

}

func (a *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {

}
