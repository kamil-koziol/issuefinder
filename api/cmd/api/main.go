package main

import (
	"log"

	"github.com/kamil-koziol/issuefinder/api/internal/server"
)

func main() {
	s := server.NewServer()
	err := s.Run()
	if err != nil {
		log.Fatal(err)
	}
}
