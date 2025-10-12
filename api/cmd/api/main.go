package main

import (
	"fmt"
	"log"

	"github.com/kamil-koziol/issuefinder/api/internal/config"
	"github.com/kamil-koziol/issuefinder/api/internal/server"
)

func main() {
	cfg := config.Config{}
	if err := cfg.Load(); err != nil {
		log.Fatal(fmt.Errorf("unable to load config: %w", err))
	}

	s, err := server.NewServer(cfg)
	if err != nil {
		log.Fatal(fmt.Errorf("unable to create server: %w", err))
	}

	err = s.Run()
	if err != nil {
		log.Fatal(err)
	}
}
