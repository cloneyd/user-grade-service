package natsstreaming

import (
	"github.com/nats-io/stan.go"
	"wb-test-task-2022/internal/config"
)

func NewStanConn(stanCfg config.StanConfig) (stan.Conn, error) {
	conn, err := stan.Connect(stanCfg.ClusterId, stanCfg.ClientId, stan.NatsURL("nats://nats-streaming:4222"))
	if err != nil {
		return nil, err
	}
	return conn, nil
}
