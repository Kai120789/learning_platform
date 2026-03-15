package router

import (
	"github.com/go-chi/chi/v5"
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

func (a *AuthRouter) AuthRoutes(
	r chi.Router,
	h AuthHandler,
	jwtMiddleware func(http.Handler) http.Handler,
) {
	r.Route("/api/auth", func(r chi.Router) {
		r.Post("/login", h.Login)
		r.Post("/register", h.Register)
		r.With(jwtMiddleware).Delete("/logout", h.Logout)
		r.With(jwtMiddleware).Post("/refresh", h.RefreshTokens)           // TODO: пока для тестов тут, потом нельзя будет вызывать рестом
		r.With(jwtMiddleware).Delete("/logout-all/{userId}", h.LogoutAll) // TODO: добавить проверку роли на админа
	})
}
