package main

import (
	"log"

	"github.com/railgodev/denet-test/internal/app"
	"github.com/railgodev/denet-test/internal/config"
)

func main() {
	// Configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// Run
	app.Run(cfg)
}
