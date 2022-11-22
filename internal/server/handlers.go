package server

import (
	"net/http"

	"github.com/gorilla/mux"

	"wb-test-task-2022/internal/middleware"
	"wb-test-task-2022/internal/usergrade/delivery/http/v1"
	"wb-test-task-2022/internal/usergrade/repository/storage"
	"wb-test-task-2022/internal/usergrade/usecase"
)

func (s *Server) MapHandlers() {
	userGradeRepo := storage.NewUserGradeRepo(s.datasource)

	userGradeUseCase := usecase.NewUserGradeUseCase(userGradeRepo)

	userGradeHandlers := v1.NewUserGradeHandlers(userGradeUseCase)

	privateRouter := mux.NewRouter().StrictSlash(true)
	publicRouter := mux.NewRouter().StrictSlash(true)

	privateRouter.Use(middleware.BasicAuth)
	privateRouter.HandleFunc("/set", userGradeHandlers.Set).Methods(http.MethodPost)
	publicRouter.HandleFunc("/get", userGradeHandlers.Get).Methods(http.MethodGet)

	s.private.Handler = privateRouter
	s.public.Handler = publicRouter
}
