package router

import (
	"github.com/go-chi/chi/v5"
	"learning-platform/api-gateway/internal/config"
	"learning-platform/api-gateway/internal/middleware"
	"learning-platform/api-gateway/internal/redis"
	"net/http"
)

type AuthRouter struct{}

type AuthHandler interface {
	Login(w http.ResponseWriter, r *http.Request)
	Register(w http.ResponseWriter, r *http.Request)
	Logout(w http.ResponseWriter, r *http.Request)
	RefreshTokens(w http.ResponseWriter, r *http.Request)
	LogoutAll(w http.ResponseWriter, r *http.Request)
}

func NewAuthRouter() *AuthRouter {
	return &AuthRouter{}
}

func (a *AuthRouter) AuthRoutes(r chi.Router, h AuthHandler, cfg *config.Config, redis *redis.RedisStorage) {
	r.Route("/api/auth", func(r chi.Router) {
		r.Post("/login", h.Login)
		r.Post("/register", h.Register)
		r.With(middleware.JWT([]byte(cfg.SignedKey), redis)).Delete("/logout", h.Logout)
		r.With(middleware.JWT([]byte(cfg.SignedKey), redis)).Post("/refresh", h.RefreshTokens)           // TODO: пока для тестов тут, потом нельзя будет вызывать рестом
		r.With(middleware.JWT([]byte(cfg.SignedKey), redis)).Delete("/logout-all/{userId}", h.LogoutAll) // TODO: добавить проверку роли на админа
	})
}
