package app

import (
	"context"
	"log/slog"
	"review/internal/config"
)

type App struct {
	config *config.Config
	log    *slog.Logger
}

func New(ctx context.Context, cfg *config.Config, log *slog.Logger) *App {

	return &App{
		config: cfg,
		log:    log,
	}
}

func (a *App) Run() error {
	return nil
}
