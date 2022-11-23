package natsstreaming

import (
	"encoding/json"
	"github.com/nats-io/stan.go"
	"log"
	"time"
	"wb-test-task-2022/internal/domain"
	"wb-test-task-2022/internal/logger"

	"wb-test-task-2022/internal/config"
)

type UserGradeSubscriber struct {
	cfg     *config.Config
	conn    stan.Conn
	handler stan.MsgHandler
}

func NewUserGradeSubscriber(cfg *config.Config, conn stan.Conn, handler stan.MsgHandler) *UserGradeSubscriber {
	return &UserGradeSubscriber{cfg: cfg, conn: conn, handler: handler}
}

func (sub *UserGradeSubscriber) Subscribe() (stan.Subscription, error) {
	return sub.conn.Subscribe(sub.cfg.StanConn.Subject, sub.handler, stan.SetManualAckMode(), stan.StartAtTime(time.Now()))
}

func HandleUserGrade(msg *stan.Msg) {
	var userGrade domain.UserGrade

	if err := msg.Ack(); err != nil {
		log.Println(err)
		return
	}

	if err := json.Unmarshal(msg.Data, &userGrade); err != nil {
		log.Printf("error unmarshalling incoming user grade JSON: %v\n", err)
		return
	}

	logger.LogUserGrade(userGrade)
}
