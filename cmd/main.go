package main

import (
	"log"

	"github.com/Semaffor/go__innotaxi_service_user/pkg/app"
)

const configsDir = "pkg/configurator/fixture"

func main() {
	if err := app.Run(configsDir); err != nil {
		log.Fatalf("%v\n", err)
	}
}
