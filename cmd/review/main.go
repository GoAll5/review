package main

import (
	"context"
	"fmt"
	"os"
	"review/internal/app"
	"review/internal/config"
	"review/internal/lib/logger"
)

// @title User Service API
// @version 1.0
// @description Review service: CRUD operations for reviews.

// @contact.name API Support Rinat
// @contact.url https://t.me/rinatgabr
// @contact.email iam.rinat.gabdrashitov@gmail.com

// @BasePath /
// @schemes http https

func main() {
	if err := run(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "error %v\n", err) // err
		os.Exit(1)
	}
}

func run() error {
	cfg := config.MustLoad()
	log := logger.New(cfg.Env)
	application := app.New(context.Background(), cfg, log) // timer
	_ = application.Run()
	return fmt.Errorf("12345")
}
