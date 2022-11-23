package main

import (
	_uuid "github.com/google/uuid"
	"log"
	"sync"
	"wb-test-task-2022/internal/config"
	"wb-test-task-2022/internal/usergrade/delivery/natsstreaming"

	"wb-test-task-2022/internal/server"
)

func main() {
	log.Println("starting service")

	cfgFile, err := config.LoadConfig("/config/config-docker")
	if err != nil {
		log.Fatalf("error loading config: %v\n", err)
	}

	cfg, err := config.ParseConfig(cfgFile)
	if err != nil {
		log.Fatalf("error parsing config: %v\n", err)
	}

	sc, err := natsstreaming.NewStanConn(cfg.StanConn)
	if err != nil {
		log.Fatalf("error connecting to nats-streamin: %v\n", err)
	}

	datasource := sync.Map{}

	uuid, err := _uuid.NewUUID()
	if err != nil {
		log.Fatalln(err)
	}

	s := server.NewServer(uuid, cfg, sc, &datasource)

	if err1, err2 := s.Run(); err1 != nil {
		log.Fatalln(err1)
	} else if err2 != nil {
		log.Fatalln(err2)
	}
}
