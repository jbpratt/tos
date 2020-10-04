package main

import (
	"log"

	"github.com/jbpratt/tos/pkg/server"
)

func main() {
	server, err := server.NewServer()
	if err != nil {
		log.Fatal(err)
	}

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
