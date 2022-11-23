package natsstreaming

import (
	"encoding/json"
	"os"

	"github.com/nats-io/stan.go"

	"wb-test-task-2022/internal/config"
	"wb-test-task-2022/internal/domain"
)

type UserGradePublisher struct {
	id   string
	cfg  *config.Config
	conn stan.Conn
}

func NewUserGradePublisher(cfg *config.Config, conn stan.Conn) *UserGradePublisher {
	return &UserGradePublisher{id: os.Getenv("REPLICA_TYPE"), cfg: cfg, conn: conn}
}

func (pub *UserGradePublisher) Publish(userGrade *domain.UserGrade) error {
	message := NewUserGradeMessage(pub.id, userGrade)

	body, err := json.Marshal(message)
	if err != nil {
		return err
	}

	return pub.conn.Publish(pub.cfg.StanConn.Subject, body)
}
