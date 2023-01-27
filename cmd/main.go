package main

import (
	"log"

	"github.com/Semaffor/go__innotaxi_service_user/pkg/server"
)

const configsDir = "pkg/config/fixture"

func main() {
	if err := server.Run(configsDir); err != nil {
		log.Fatalf("%v\n", err)
	}
}
