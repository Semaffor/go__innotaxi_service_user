package main

import (
	"log"

	"github.com/Semaffor/go__innotaxi_service_user/pkg/app"
)

func main() {
	if err := app.Run(); err != nil {
		log.Fatalf("%v\n", err)
	}
}
