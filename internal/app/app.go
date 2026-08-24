package app

import (
	"context"
	"log/slog"
	"review/internal/config"
	"review/internal/transport/http/echoserver"
)

type App struct {
	HTTPSrv *echoserver.Server
	config  *config.Config
	log     *slog.Logger
}

func New(ctx context.Context, cfg *config.Config, log *slog.Logger) *App {

	httpSrv, _ := setupHttp(ctx, log, cfg.HTTP.Addr, nil)

	return &App{
		config:  cfg,
		log:     log,
		HTTPSrv: httpSrv,
	}
}

func (a *App) Run() error {
	return nil
}
