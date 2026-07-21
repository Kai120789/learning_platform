package router

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

type GroupRouter struct {
}

type GroupHandler interface {
	CreateGroup(w http.ResponseWriter, r *http.Request)
	UpdateGroup(w http.ResponseWriter, r *http.Request)
	RemoveGroup(w http.ResponseWriter, r *http.Request)
	GetGroupById(w http.ResponseWriter, r *http.Request)
	GetGroups(w http.ResponseWriter, r *http.Request)
	AddUsersToGroup(w http.ResponseWriter, r *http.Request)
	RemoveUserFromGroup(w http.ResponseWriter, r *http.Request)
	GetUserGroups(w http.ResponseWriter, r *http.Request)
	GetGroupsByTutorId(w http.ResponseWriter, r *http.Request)
	GetGroupUsers(w http.ResponseWriter, r *http.Request)
}

func NewGroupRouter() *GroupRouter {
	return &GroupRouter{}
}

func (u *GroupRouter) GroupRoutes(
	r chi.Router,
	h GroupHandler,
	jwtMiddleware func(http.Handler) http.Handler,
) {
	r.With(jwtMiddleware).Route("/api/group", func(r chi.Router) {
		r.Post("/", h.CreateGroup)
		r.Patch("/{groupId}", h.UpdateGroup)
		r.Delete("/{groupId}", h.RemoveGroup)
		r.Get("/{groupId}", h.GetGroupById)
		r.Get("/", h.GetGroups)
		r.Post("/{groupId}/add-user", h.AddUsersToGroup)
		r.Delete("/{groupId}/remove-user/{userId}", h.RemoveUserFromGroup)
		r.Get("/user", h.GetUserGroups)
		r.Get("/tutor", h.GetGroupsByTutorId)
		r.Get("/{groupId}/get-users", h.GetGroupUsers)
	})
}
