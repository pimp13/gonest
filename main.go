package main

import (
	"github.com/pimp13/gonest/app"
	"github.com/pimp13/gonest/common/database"
	"github.com/pimp13/gonest/config"
	"log"
)

func main() {
	cfg := config.LoadConfig()

	// Initialize database connection
	database.Init()

	application := app.NewApp(cfg)

	if err := application.Bootstrap(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
