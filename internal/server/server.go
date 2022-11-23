package server

import (
	"context"
	"github.com/nats-io/stan.go"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
	"wb-test-task-2022/internal/config"
	"wb-test-task-2022/internal/usergrade/delivery/natsstreaming"
)

const ctxTimeout = 5

type Server struct {
	cfg        *config.Config
	public     *http.Server
	private    *http.Server
	conn       stan.Conn
	datasource *sync.Map
}

func NewServer(cfg *config.Config, conn stan.Conn, datasource *sync.Map) *Server {
	return &Server{
		cfg:        cfg,
		public:     &http.Server{Addr: cfg.PublicServer.Addr},
		private:    &http.Server{Addr: cfg.PrivateServer.Addr},
		conn:       conn,
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

	var sub stan.Subscription
	go func(sub stan.Subscription) {
		subscriber := natsstreaming.NewUserGradeSubscriber(s.cfg, s.conn, natsstreaming.HandleUserGrade)
		sub, err := subscriber.Subscribe()
		if err != nil {
			log.Fatalln(err)
		}
	}(sub)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit

	if err := sub.Unsubscribe(); err != nil {
		log.Fatalln(err)
	}

	if err := sub.Close(); err != nil {
		log.Fatalln(err)
	}

	ctx, shutdown := context.WithTimeout(context.Background(), ctxTimeout*time.Second)
	defer shutdown()

	log.Println("server gracefully shut down")
	return s.private.Shutdown(ctx), s.public.Shutdown(ctx)
}
