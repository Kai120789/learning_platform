package router

import (
	"github.com/go-chi/chi/v5"
	"learning-platform/api-gateway/internal/dto/enum"
	"net/http"
)

type BoardRouter struct{}

type BoardHandler interface {
	GetBoard(w http.ResponseWriter, r *http.Request)
	CreateBoard(w http.ResponseWriter, r *http.Request)
	GetTutorBoards(w http.ResponseWriter, r *http.Request)
	DeleteBoard(w http.ResponseWriter, r *http.Request)
	UpdateBoard(w http.ResponseWriter, r *http.Request)
	ProxyBoardWS(w http.ResponseWriter, r *http.Request)
}

func NewBoardRouter() *BoardRouter {
	return &BoardRouter{}
}

func (u *BoardRouter) BoardRoutes(
	r chi.Router,
	h BoardHandler,
	jwtMiddleware func(http.Handler) http.Handler,
	roleMiddleware func(minNeededRole enum.UserRole) func(http.Handler) http.Handler,
) {
	r.With(jwtMiddleware).Route("/api/board", func(r chi.Router) {
		r.With(roleMiddleware(enum.RoleStudent)).Get("/{boardID}", h.GetBoard)
		r.With(roleMiddleware(enum.RoleTutor)).Post("/", h.CreateBoard)
		r.With(roleMiddleware(enum.RoleTutor)).Get("/", h.GetTutorBoards)
		r.With(roleMiddleware(enum.RoleTutor)).Delete("/{boardID}", h.DeleteBoard)
		r.With(roleMiddleware(enum.RoleTutor)).Put("/{boardID}", h.UpdateBoard)

		r.With(roleMiddleware(enum.RoleStudent)).Get("/{boardID}/ws", h.ProxyBoardWS)
	})
}
