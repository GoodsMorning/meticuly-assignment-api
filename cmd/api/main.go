package main

import (
	"log"
	"meticuly-assignment-api/internal/server"
)

func main() {

	if err := server.RunServer(); err != nil {
		log.Fatal("Error cannot run server")
	}
}