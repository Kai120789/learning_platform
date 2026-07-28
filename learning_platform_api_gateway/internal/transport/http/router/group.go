package router

import (
	"github.com/go-chi/chi/v5"
	"learning-platform/api-gateway/internal/dto/enum"
	"net/http"
)

type GroupRouter struct {
}

type GroupHandler interface {
	CreateGroup(w http.ResponseWriter, r *http.Request)
	UpdateGroup(w http.ResponseWriter, r *http.Request)
	RemoveGroup(w http.ResponseWriter, r *http.Request)
	GetGroupByID(w http.ResponseWriter, r *http.Request)
	GetGroups(w http.ResponseWriter, r *http.Request)
	AddUsersToGroup(w http.ResponseWriter, r *http.Request)
	RemoveUserFromGroup(w http.ResponseWriter, r *http.Request)
	GetGroupsByStudentID(w http.ResponseWriter, r *http.Request)
	GetGroupsByTutorID(w http.ResponseWriter, r *http.Request)
	GetGroupUsers(w http.ResponseWriter, r *http.Request)
}

func NewGroupRouter() *GroupRouter {
	return &GroupRouter{}
}

func (u *GroupRouter) GroupRoutes(
	r chi.Router,
	h GroupHandler,
	jwtMiddleware func(http.Handler) http.Handler,
	roleMiddleware func(minNeededRole enum.UserRole) func(http.Handler) http.Handler,
) {
	r.With(jwtMiddleware).Route("/api/group", func(r chi.Router) {
		r.With(roleMiddleware(enum.RoleTutor)).Post("/", h.CreateGroup)
		r.With(roleMiddleware(enum.RoleTutor)).Patch("/{groupID}", h.UpdateGroup)
		r.With(roleMiddleware(enum.RoleTutor)).Delete("/{groupID}", h.RemoveGroup)
		r.With(roleMiddleware(enum.RoleStudent)).Get("/{groupID}", h.GetGroupByID)
		r.With(roleMiddleware(enum.RoleStudent)).Get("/", h.GetGroups)
		r.With(roleMiddleware(enum.RoleTutor)).Post("/{groupID}/add-user", h.AddUsersToGroup)
		r.With(roleMiddleware(enum.RoleTutor)).Delete("/{groupID}/remove-user/{userID}", h.RemoveUserFromGroup)
		r.With(roleMiddleware(enum.RoleStudent)).Get("/student", h.GetGroupsByStudentID)
		r.With(roleMiddleware(enum.RoleTutor)).Get("/tutor", h.GetGroupsByTutorID)
		r.With(roleMiddleware(enum.RoleStudent)).Get("/{groupID}/get-users", h.GetGroupUsers)
	})
}
