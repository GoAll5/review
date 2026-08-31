package app

import (
	"context"
	"log/slog"
	"review/internal/config"
	"review/internal/store/mongostore"
	"review/internal/transport/http/echoserver"
)

type App struct {
	ClientDB *mongostore.DB
	HTTPSrv  *echoserver.Server
	config   *config.Config
	log      *slog.Logger
}

func New(ctx context.Context, cfg *config.Config, log *slog.Logger) (*App, error) {
	infra, err := initInfra(ctx, cfg.Mongo.DSN)
	if err != nil {
		return nil, err
	}

	httpSrv, err := setupHttp(ctx, log, cfg.HTTP.Addr, nil)
	if err != nil {
		return nil, err
	}

	return &App{
		config:   cfg,
		ClientDB: infra.ClientDB,
		log:      log,
		HTTPSrv:  httpSrv,
	}, nil
}

func (a *App) Run() error {
	return nil
}
