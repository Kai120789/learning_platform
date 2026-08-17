package router

import (
	"github.com/go-chi/chi/v5"
	"learning-platform/api-gateway/internal/dto/enum"
	"net/http"
)

type MaterialRouter struct {
}

type MaterialHandler interface {
	CreateFolder(w http.ResponseWriter, r *http.Request)
	MoveFolders(w http.ResponseWriter, r *http.Request)
	RenameFolder(w http.ResponseWriter, r *http.Request)
	DeleteOneFolder(w http.ResponseWriter, r *http.Request)
	DeleteFolders(w http.ResponseWriter, r *http.Request)
	GetMaterials(w http.ResponseWriter, r *http.Request)
	CreateMaterials(w http.ResponseWriter, r *http.Request)
	MoveMaterials(w http.ResponseWriter, r *http.Request)
	RenameMaterial(w http.ResponseWriter, r *http.Request)
	DeleteOneMaterial(w http.ResponseWriter, r *http.Request)
	DeleteMaterials(w http.ResponseWriter, r *http.Request)
	UpdateUsersMaterialsAccess(w http.ResponseWriter, r *http.Request)
	UpdateUsersFoldersAccess(w http.ResponseWriter, r *http.Request)
}

func NewMaterialRouter() *MaterialRouter {
	return &MaterialRouter{}
}

func (l *MaterialRouter) MaterialRoutes(
	r chi.Router,
	h MaterialHandler,
	jwtMiddleware func(http.Handler) http.Handler,
	roleMiddleware func(minNeededRole enum.UserRole) func(http.Handler) http.Handler,
) {
	r.With(jwtMiddleware).Route("/api/material", func(r chi.Router) {
		r.With(roleMiddleware(enum.RoleTutor)).Post("/folder", h.CreateFolder)
		r.With(roleMiddleware(enum.RoleTutor)).Patch("/folder", h.MoveFolders)
		r.With(roleMiddleware(enum.RoleTutor)).Patch("/folder/{folderID}", h.RenameFolder)
		r.With(roleMiddleware(enum.RoleTutor)).Delete("/folder/{folderID}", h.DeleteOneFolder)
		r.With(roleMiddleware(enum.RoleTutor)).Delete("/folder", h.DeleteFolders)
		r.With(roleMiddleware(enum.RoleStudent)).Get("/{folderID}", h.GetMaterials)
		r.With(roleMiddleware(enum.RoleTutor)).Post("/{folderID}", h.CreateMaterials)
		r.With(roleMiddleware(enum.RoleTutor)).Patch("/", h.MoveMaterials)
		r.With(roleMiddleware(enum.RoleTutor)).Patch("/{materialID}", h.RenameMaterial)
		r.With(roleMiddleware(enum.RoleTutor)).Delete("/{materialID}", h.DeleteOneMaterial)
		r.With(roleMiddleware(enum.RoleTutor)).Delete("/", h.DeleteMaterials)
		r.With(roleMiddleware(enum.RoleTutor)).Put("/", h.UpdateUsersMaterialsAccess)
		r.With(roleMiddleware(enum.RoleTutor)).Put("/folder", h.UpdateUsersFoldersAccess)
	})
}
