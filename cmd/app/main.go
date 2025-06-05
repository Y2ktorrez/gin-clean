package main

import (
	"fmt"
	"gin-clean/config"
	"gin-clean/internal/app"
	"log"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	app, err := app.New(config)
	if err != nil {
		log.Fatalf("Failed to create app: %v", err)
	}

	addr := fmt.Sprintf(":%s", config.AppPort)
	log.Printf("Server starting on %s", addr)
	if err := app.Run(addr); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
