package router

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

type UserRouter struct{}

type UserHandler interface {
	GetUserById(w http.ResponseWriter, r *http.Request)
	GetUserData(w http.ResponseWriter, r *http.Request)
	CreateUser(w http.ResponseWriter, r *http.Request)
}

func NewUserRouter() *UserRouter {
	return &UserRouter{}
}

func (u *UserRouter) UserRoutes(
	r chi.Router,
	h UserHandler,
	jwtMiddleware func(http.Handler) http.Handler,
) {
	r.With(jwtMiddleware).Route("/api/user", func(r chi.Router) {
		r.Get("/{userId}", h.GetUserById)
		r.Get("/data", h.GetUserData)
		r.Post("/", h.CreateUser)
	})
}
