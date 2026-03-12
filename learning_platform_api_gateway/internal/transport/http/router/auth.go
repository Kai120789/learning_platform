package router

import (
	"github.com/go-chi/chi/v5"
	"learning-platform/api-gateway/internal/config"
	"learning-platform/api-gateway/internal/middleware"
	"net/http"
)

type AuthRouter struct{}

type AuthHandler interface {
	Login(w http.ResponseWriter, r *http.Request)
	Register(w http.ResponseWriter, r *http.Request)
	Logout(w http.ResponseWriter, r *http.Request)
}

func NewAuthRouter() *AuthRouter {
	return &AuthRouter{}
}

func (a *AuthRouter) AuthRoutes(r chi.Router, h AuthHandler, cfg *config.Config) {
	r.Route("/api/auth", func(r chi.Router) {
		r.Post("/login", h.Login)
		r.Post("/register", h.Register)
		r.With(middleware.JWT([]byte(cfg.SignedKey))).Delete("/logout", h.Logout)
	})
}
