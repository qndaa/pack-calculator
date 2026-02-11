package main

import (
	"log"

	"github.com/qndaa/pack-calculator/internal/app"
)

func main() {
	server, err := app.New()
	if err != nil {
		log.Fatalf("failed to create app: %s", err.Error())
	}

	if err := server.Run(); err != nil {
		log.Fatalf("app.Run() exit with error: %s", err.Error())
	}
}
