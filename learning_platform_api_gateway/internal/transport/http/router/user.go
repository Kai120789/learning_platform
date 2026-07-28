package router

import (
	"github.com/go-chi/chi/v5"
	"learning-platform/api-gateway/internal/dto/enum"
	"net/http"
)

type UserRouter struct{}

type UserHandler interface {
	GetUserByID(w http.ResponseWriter, r *http.Request)
	GetUserData(w http.ResponseWriter, r *http.Request)
	CreateUser(w http.ResponseWriter, r *http.Request)
	UpdateUserInfo(w http.ResponseWriter, r *http.Request)
	UpdateUserSettings(w http.ResponseWriter, r *http.Request)
	UpdateUserTheme(w http.ResponseWriter, r *http.Request)
	UpdateUserAvatar(w http.ResponseWriter, r *http.Request)
	UpdateUserTgUsername(w http.ResponseWriter, r *http.Request)
	ChangeUserEmail(w http.ResponseWriter, r *http.Request)
	ChangeUserPassword(w http.ResponseWriter, r *http.Request)
	GetUsersWithPagination(w http.ResponseWriter, r *http.Request)
}

func NewUserRouter() *UserRouter {
	return &UserRouter{}
}

func (u *UserRouter) UserRoutes(
	r chi.Router,
	h UserHandler,
	jwtMiddleware func(http.Handler) http.Handler,
	roleMiddleware func(minNeededRole enum.UserRole) func(http.Handler) http.Handler,
) {
	r.With(jwtMiddleware).Route("/api/user", func(r chi.Router) {
		r.With(roleMiddleware(enum.RoleStudent)).Get("/{userID}", h.GetUserByID)
		r.With(roleMiddleware(enum.RoleStudent)).Get("/data", h.GetUserData)
		r.With(roleMiddleware(enum.RoleStudent)).Put("/info", h.UpdateUserInfo)
		r.With(roleMiddleware(enum.RoleStudent)).Put("/settings", h.UpdateUserSettings)
		r.With(roleMiddleware(enum.RoleStudent)).Patch("/theme", h.UpdateUserTheme)
		r.With(roleMiddleware(enum.RoleStudent)).Patch("/avatar", h.UpdateUserAvatar)
		r.With(roleMiddleware(enum.RoleStudent)).Patch("/tg", h.UpdateUserTgUsername)
		r.With(roleMiddleware(enum.RoleStudent)).Patch("/email", h.ChangeUserEmail)
		r.With(roleMiddleware(enum.RoleStudent)).Patch("/password", h.ChangeUserPassword)
		r.With(roleMiddleware(enum.RoleAdmin)).Post("/", h.CreateUser)
		r.With(roleMiddleware(enum.RoleTutor)).Get("/", h.GetUsersWithPagination)
	})
}
