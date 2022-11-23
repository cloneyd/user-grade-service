package natsstreaming

import (
	"encoding/json"
	"log"
	"time"

	_uuid "github.com/google/uuid"
	"github.com/nats-io/stan.go"

	"wb-test-task-2022/internal/config"
	"wb-test-task-2022/internal/logger"
)

type UserGradeSubscriber struct {
	uuid    _uuid.UUID
	cfg     *config.Config
	conn    stan.Conn
	handler stan.MsgHandler
}

func NewUserGradeSubscriber(uuid _uuid.UUID, cfg *config.Config, conn stan.Conn, handler stan.MsgHandler) *UserGradeSubscriber {
	return &UserGradeSubscriber{uuid: uuid, cfg: cfg, conn: conn, handler: handler}
}

func (sub *UserGradeSubscriber) Subscribe(time time.Time) (stan.Subscription, error) {
	return sub.conn.Subscribe(sub.cfg.StanConn.Subject, sub.handler, stan.SetManualAckMode(), stan.StartAtTime(time))
}

func HandleUserGrade(msg *stan.Msg) {
	var userGradeMessage *UserGradeMessage

	if err := msg.Ack(); err != nil {
		log.Println(err)
		return
	}

	if err := json.Unmarshal(msg.Data, &userGradeMessage); err != nil {
		log.Printf("error unmarshalling incoming user grade JSON: %v\n", err)
		return
	}

	//if userGradeMessage.PublisherUUID == sub.uuid {
	//	return
	//}

	logger.LogUserGrade(userGradeMessage.Payload)
}
