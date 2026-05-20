package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Router struct {
	UserRouter  UserRouter
	AuthRouter  AuthRouter
	GroupRouter GroupRouter
}

type Handler struct {
	UserHandler  UserHandler
	AuthHandler  AuthHandler
	GroupHandler GroupHandler
}

func New(
	handler *Handler,
	jwtMiddleware func(http.Handler) http.Handler,
) http.Handler {
	r := chi.NewRouter()

	router := &Router{
		AuthRouter:  *NewAuthRouter(),
		UserRouter:  *NewUserRouter(),
		GroupRouter: *NewGroupRouter(),
	}

	router.UserRouter.UserRoutes(r, handler.UserHandler, jwtMiddleware)
	router.AuthRouter.AuthRoutes(r, handler.AuthHandler, jwtMiddleware)
	router.GroupRouter.GroupRoutes(r, handler.GroupHandler, jwtMiddleware)

	return r
}
