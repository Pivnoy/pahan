package main

import (
	"log"
	"pahan/config"
	"pahan/internal/app"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Error in parse config: %s\n", err)
	}

	app.Run(cfg)
}
