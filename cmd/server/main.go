package main

import (
	"log"
	"sync"

	"wb-test-task-2022/internal/server"
	"wb-test-task-2022/pkg/config"
)

func main() {
	log.Println("starting service")

	cfgFile, err := config.LoadConfig("/config/config-docker")
	if err != nil {
		log.Fatalf("error loading config: %v", err)
	}

	cfg, err := config.ParseConfig(cfgFile)
	if err != nil {
		log.Fatalf("error parsing config: %v", err)
	}

	datasource := sync.Map{}
	s := server.NewServer(cfg, &datasource)

	if err1, err2 := s.Run(); err1 != nil {
		log.Fatalln(err1)
	} else if err2 != nil {
		log.Fatalln(err2)
	}
}
