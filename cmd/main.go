package main

import (
	"log"

	"github.com/qndaa/pack-calculator/internal/app"
)

func main() {
	server := app.New()
	if err := server.Run(); err != nil {
		log.Fatalf("app.Run() exit with error: %s", err.Error())
	}
}
