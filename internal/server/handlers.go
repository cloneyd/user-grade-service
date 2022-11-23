package server

import (
	"net/http"
	"wb-test-task-2022/internal/usergrade/delivery/natsstreaming"

	"github.com/gorilla/mux"

	"wb-test-task-2022/internal/middleware"
	"wb-test-task-2022/internal/usergrade/delivery/http/v1"
	"wb-test-task-2022/internal/usergrade/repository/storage"
	"wb-test-task-2022/internal/usergrade/usecase"
)

func (s *Server) MapHandlers() {
	userGradeRepo := storage.NewUserGradeRepo(s.datasource)

	userGradePublisher := natsstreaming.NewUserGradePublisher(s.cfg, s.conn)

	userGradeUseCase := usecase.NewUserGradeUseCase(userGradePublisher, userGradeRepo)

	userGradeHandlers := v1.NewUserGradeHandlers(userGradeUseCase)

	privateRouter := mux.NewRouter().StrictSlash(true)
	privateRouter.Use(middleware.BasicAuth)
	privateRouter.HandleFunc("/set", userGradeHandlers.Set).Methods(http.MethodPost)

	publicRouter := mux.NewRouter().StrictSlash(true)
	publicRouter.HandleFunc("/get", userGradeHandlers.Get).Methods(http.MethodGet)
	publicRouter.HandleFunc("/backup", userGradeHandlers.Backup).Methods(http.MethodGet) // нужна аутентификация

	s.private.Handler = privateRouter
	s.public.Handler = publicRouter
}
