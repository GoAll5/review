package app

import (
	"context"
	"log/slog"
	"review/internal/config"
)

type infra struct {
}

func initInfra(ctx context.Context, cfg *config.Config, log *slog.Logger) (*infra, error) {
	return nil, nil
}
