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
	UpdateUserInfo(w http.ResponseWriter, r *http.Request)
	UpdateUserSettings(w http.ResponseWriter, r *http.Request)
	UpdateUserTheme(w http.ResponseWriter, r *http.Request)
	UpdateUserAvatar(w http.ResponseWriter, r *http.Request)
	UpdateUserTgUsername(w http.ResponseWriter, r *http.Request)
	ChangeUserEmail(w http.ResponseWriter, r *http.Request)
	ChangeUserPassword(w http.ResponseWriter, r *http.Request)
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
		r.Put("/info", h.UpdateUserInfo)
		r.Put("/settings", h.UpdateUserSettings)
		r.Patch("/theme", h.UpdateUserTheme)
		r.Patch("/avatar", h.UpdateUserAvatar)
		r.Patch("/tg", h.UpdateUserTgUsername)
		r.Patch("/email", h.ChangeUserEmail)
		r.Patch("/password", h.ChangeUserPassword)
		r.Post("/", h.CreateUser)
	})
}
