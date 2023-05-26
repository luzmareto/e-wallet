package main

import (
	"log"

	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-5/khilmi-aminudin/challenge/go-ewallet/api"
	"git.enigmacamp.com/enigma-camp/enigmacamp-2.0/batch-5/khilmi-aminudin/challenge/go-ewallet/utils"
)

func main() {
	cfg, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	server, err := api.NewServer(*cfg)
	if err != nil {
		log.Fatalf("Failed to create server: %v", err)
	}
	log.Fatal(server.Start(cfg.ServerAddress))
}
