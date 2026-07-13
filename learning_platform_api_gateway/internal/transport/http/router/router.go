package router

import (
	"github.com/go-chi/cors"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Router struct {
	UserRouter     *UserRouter
	AuthRouter     *AuthRouter
	GroupRouter    *GroupRouter
	LessonRouter   *LessonRouter
	ScheduleRouter *ScheduleRouter
	SubjectRouter  *SubjectRouter
}

type Handler struct {
	UserHandler     UserHandler
	AuthHandler     AuthHandler
	GroupHandler    GroupHandler
	LessonHandler   LessonHandler
	ScheduleHandler ScheduleHandler
	SubjectHandler  SubjectHandler
}

func New(
	handler *Handler,
	jwtMiddleware func(http.Handler) http.Handler,
) http.Handler {
	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{
			"http://localhost:5173",
		},
		AllowedMethods: []string{
			"GET",
			"POST",
			"PUT",
			"DELETE",
			"OPTIONS",
		},
		AllowedHeaders: []string{
			"Accept",
			"Authorization",
			"Content-Type",
		},
		AllowCredentials: true,
	}))

	router := &Router{
		AuthRouter:     NewAuthRouter(),
		UserRouter:     NewUserRouter(),
		GroupRouter:    NewGroupRouter(),
		LessonRouter:   NewLessonRouter(),
		ScheduleRouter: NewScheduleRouter(),
		SubjectRouter:  NewSubjectRouter(),
	}

	router.UserRouter.UserRoutes(r, handler.UserHandler, jwtMiddleware)
	router.AuthRouter.AuthRoutes(r, handler.AuthHandler, jwtMiddleware)
	router.GroupRouter.GroupRoutes(r, handler.GroupHandler, jwtMiddleware)
	router.LessonRouter.LessonRoutes(r, handler.LessonHandler, jwtMiddleware)
	router.ScheduleRouter.ScheduleRoutes(r, handler.ScheduleHandler, jwtMiddleware)
	router.SubjectRouter.SubjectRoutes(r, handler.SubjectHandler, jwtMiddleware)

	return r
}
