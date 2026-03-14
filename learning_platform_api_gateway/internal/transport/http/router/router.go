package router

import (
	"github.com/go-chi/chi/v5"
	"learning-platform/api-gateway/internal/config"
	"learning-platform/api-gateway/internal/redis"
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

func New(handler *Handler, cfg *config.Config, redis *redis.RedisStorage) http.Handler {
	r := chi.NewRouter()

	router := &Router{
		AuthRouter: *NewAuthRouter(),
		UserRouter: *NewUserRouter(),
	}

	router.UserRouter.UserRoutes(r, handler.UserHandler)
	router.AuthRouter.AuthRoutes(r, handler.AuthHandler, cfg, redis)

	return r
}
