package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"wb-test-task-2022/pkg/config"
)

const ctxTimeout = 5

type Server struct {
	public     *http.Server
	private    *http.Server
	datasource *sync.Map
}

func NewServer(cfg *config.Config, datasource *sync.Map) *Server {
	return &Server{
		public:     &http.Server{Addr: cfg.PublicServer.Addr},
		private:    &http.Server{Addr: cfg.PrivateServer.Addr},
		datasource: datasource,
	}
}

func (s *Server) Run() (error, error) {
	s.MapHandlers()

	go func() {
		if err := s.private.ListenAndServe(); err != nil {
			log.Fatalln(err)
		}
	}()

	go func() {
		if err := s.public.ListenAndServe(); err != nil {
			log.Fatalln(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), ctxTimeout*time.Second)
	defer shutdown()

	log.Println("server gracefully shut down")
	return s.private.Shutdown(ctx), s.public.Shutdown(ctx)
}
