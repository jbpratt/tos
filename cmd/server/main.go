package main

import (
	"log"

	"github.com/jbpratt/tos/internal/server"
	_ "github.com/mattn/go-sqlite3" // db
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
