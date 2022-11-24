package natsstreaming

import (
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/nats-io/stan.go"

	"wb-test-task-2022/internal/config"
	"wb-test-task-2022/internal/logger"
)

type UserGradeSubscriber struct {
	id        string
	cfg       *config.Config
	conn      stan.Conn
	timestamp time.Time
	handler   stan.MsgHandler
}

func NewUserGradeSubscriber(cfg *config.Config, conn stan.Conn, timestamp time.Time, handler stan.MsgHandler) *UserGradeSubscriber {
	return &UserGradeSubscriber{id: os.Getenv("REPLICA_TYPE"), cfg: cfg, conn: conn, timestamp: timestamp, handler: handler}
}

func (sub *UserGradeSubscriber) Subscribe() (stan.Subscription, error) {
	return sub.conn.Subscribe(sub.cfg.StanConn.Subject, sub.handler, stan.SetManualAckMode(), stan.StartAtTime(sub.timestamp))
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

	// Проверяем PublisherId, чтобы не обрабатывать свои сообщения
	if userGradeMessage.PublisherId == os.Getenv("REPLICA_TYPE") {
		return
	}

	//TODO: Здесь должно быть сохранение данных других реплик в storage, но я не успел
	logger.LogUserGrade(userGradeMessage.Payload)
}
