package app

import (
	"context"
	"review/internal/store/mongostore"
)

type infra struct {
	ClientDB *mongostore.DB
}

// log *slog.Logger
func initInfra(ctx context.Context, dsn string) (*infra, error) {
	clientDB, err := mongostore.New(ctx, dsn)
	if err != nil {
		return nil, err
	}

	return &infra{ClientDB: clientDB}, nil
}
