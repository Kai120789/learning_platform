package router

import (
	"github.com/go-chi/chi/v5"
)

type AuthRouter struct{}

type AuthHandler interface {
}

func NewAuthRouter() *AuthRouter {
	return &AuthRouter{}
}

func (a *AuthRouter) AuthRoutes(r chi.Router, h AuthHandler) {
	r.Route("/api/auth", func(r chi.Router) {

	})
}
