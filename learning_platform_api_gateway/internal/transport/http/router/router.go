package router

import (
	"github.com/go-chi/chi/v5"
	"learning-platform/api-gateway/internal/config"
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

func New(handler *Handler, cfg *config.Config) http.Handler {
	r := chi.NewRouter()

	router := &Router{
		AuthRouter: *NewAuthRouter(),
		UserRouter: *NewUserRouter(),
	}

	router.UserRouter.UserRoutes(r, handler.UserHandler)
	router.AuthRouter.AuthRoutes(r, handler.AuthHandler, cfg)

	return r
}
