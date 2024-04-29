package main

import (
	"log"

	"github.com/Angstreminus/exchanger/internal/server"
	"github.com/Angstreminus/exchanger/pkg/config"
	"github.com/Angstreminus/exchanger/pkg/logger"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err.Error())
	}
	log := logger.MustInitLogger(&cfg)
	server := server.NewServer(&cfg, log)
	server.MustRunWithGracefullShutdown()
}
