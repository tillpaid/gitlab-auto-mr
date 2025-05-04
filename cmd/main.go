package main

import (
	"fmt"
	"os"

	"github.com/tillpaid/gitlab-auto-mr/internal/application"
	"github.com/tillpaid/gitlab-auto-mr/internal/config"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		fmt.Println("⚠️", err)
		os.Exit(1)
	}

	app := application.New(cfg)

	if err := app.Run(); err != nil {
		fmt.Println("\n⚠️", err)
		os.Exit(1)
	}
}
