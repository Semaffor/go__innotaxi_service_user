package main

import (
	"innotaxi"
	"log"
)

func main() {
	server := new(innotaxi.Server)
	if err := server.Run("8000"); err != nil {
		log.Fatalf("error occured while running: %s", err.Error())
	}
}
