package main

import (
	"log"
	"os"
	"sync"
	"time"
	"wb-test-task-2022/internal/backup"
	"wb-test-task-2022/internal/config"
	"wb-test-task-2022/internal/usergrade/delivery/natsstreaming"

	"wb-test-task-2022/internal/server"
)

func main() {
	log.Println("starting service")

	cfgFile, err := config.LoadConfig(os.Getenv("CONFIG_PATH"))
	if err != nil {
		log.Fatalf("error loading config: %v\n", err)
	}

	cfg, err := config.ParseConfig(cfgFile)
	if err != nil {
		log.Fatalf("error parsing config: %v\n", err)
	}

	datasource := sync.Map{}

	timestamp := time.Now()
	replicaType := os.Getenv("REPLICA_TYPE")
	// 1. в текущем примере гибель мастера может быть фатальной, и вести к рассинхронизации реплик
	// 2. мастер отличается от слейвов
	if replicaType == "slave" {
		timestamp, err = backup.LoadBackup(&datasource)
		if err != nil {
			log.Fatalln(err)
		}
	}

	sc, err := natsstreaming.NewStanConn(cfg.StanConn)
	if err != nil {
		log.Fatalf("error connecting to nats-streamin: %v\n", err)
	}

	s := server.NewServer(cfg, sc, timestamp, &datasource)

	if err1, err2 := s.Run(); err1 != nil {
		log.Fatalln(err1)
	} else if err2 != nil {
		log.Fatalln(err2)
	}
}
