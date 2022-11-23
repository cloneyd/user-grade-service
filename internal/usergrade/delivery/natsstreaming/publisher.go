package natsstreaming

import (
	"encoding/json"

	_uuid "github.com/google/uuid"
	"github.com/nats-io/stan.go"

	"wb-test-task-2022/internal/config"
	"wb-test-task-2022/internal/domain"
)

type UserGradePublisher struct {
	uuid _uuid.UUID
	cfg  *config.Config
	conn stan.Conn
}

func NewUserGradePublisher(uuid _uuid.UUID, cfg *config.Config, conn stan.Conn) *UserGradePublisher {
	return &UserGradePublisher{uuid: uuid, cfg: cfg, conn: conn}
}

func (pub *UserGradePublisher) Publish(userGrade *domain.UserGrade) error {
	message := NewUserGradeMessage(pub.uuid, userGrade)

	body, err := json.Marshal(message)
	if err != nil {
		return err
	}

	return pub.conn.Publish(pub.cfg.StanConn.Subject, body)
}
