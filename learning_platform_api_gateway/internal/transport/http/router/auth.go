package router

import (
	"github.com/go-chi/chi/v5"
	"learning-platform/api-gateway/internal/dto/enum"
	"net/http"
)

type AuthRouter struct{}

type AuthHandler interface {
	Login(w http.ResponseWriter, r *http.Request)
	Register(w http.ResponseWriter, r *http.Request)
	Logout(w http.ResponseWriter, r *http.Request)
	LogoutAll(w http.ResponseWriter, r *http.Request)
}

func NewAuthRouter() *AuthRouter {
	return &AuthRouter{}
}

func (a *AuthRouter) AuthRoutes(
	r chi.Router,
	h AuthHandler,
	jwtMiddleware func(http.Handler) http.Handler,
	roleMiddleware func(minNeededRole enum.UserRole) func(http.Handler) http.Handler,
) {
	r.Route("/api/auth", func(r chi.Router) {
		r.Post("/login", h.Login)
		r.Post("/register", h.Register)
		r.With(jwtMiddleware).Delete("/logout", h.Logout)
		r.With(jwtMiddleware).With(roleMiddleware(enum.RoleAdmin)).Delete("/logout-all/{userID}", h.LogoutAll)
	})
}
