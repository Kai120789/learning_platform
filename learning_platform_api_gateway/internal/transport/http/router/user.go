package router

import "github.com/go-chi/chi/v5"

type UserRouter struct{}

type UserHandler interface {
}

func NewUserRouter() *UserRouter {
	return &UserRouter{}
}

func (u *UserRouter) UserRoutes(r chi.Router, h UserHandler) {
	r.Route("/api/user", func(r chi.Router) {

	})
}
