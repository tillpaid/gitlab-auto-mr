package main

import (
	"log"

	"github.com/tillpaid/gitlab-auto-mr/internal/application"
	"github.com/tillpaid/gitlab-auto-mr/internal/config"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("Load config failed:", err)
	}

	app := application.New(cfg)

	if err := app.Run(); err != nil {
		log.Fatal("Application failed:", err)
	}
}
