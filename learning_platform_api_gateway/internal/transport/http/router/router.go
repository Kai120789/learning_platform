package router

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

type Router struct {
	UserRouter UserRouter
	AuthRouter AuthRouter
}

type Handler struct {
	UserHandler UserHandler
	AuthHandler AuthHandler
}

func New(handler *Handler) http.Handler {
	r := chi.NewRouter()

	router := &Router{
		AuthRouter: *NewAuthRouter(),
		UserRouter: *NewUserRouter(),
	}

	router.UserRouter.UserRoutes(r, handler)
	router.AuthRouter.AuthRoutes(r, handler)

	return r
}
