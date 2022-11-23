package natsstreaming

import (
	"github.com/nats-io/stan.go"
	"wb-test-task-2022/internal/config"
)

type UserGradePublisher struct {
	cfg  *config.Config
	conn stan.Conn
}

func NewUserGradePublisher(cfg *config.Config, conn stan.Conn) *UserGradePublisher {
	return &UserGradePublisher{cfg: cfg, conn: conn}
}

func (pub *UserGradePublisher) Publish(body []byte) error {
	return pub.conn.Publish(pub.cfg.StanConn.Subject, body)
}
