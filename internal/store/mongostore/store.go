package mongostore

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"time"
)

type DB struct {
	Client *mongo.Client
}

func New(ctx context.Context, dsn string) (*DB, error) {
	op := "store.mongo.New()"

	clientOpts := options.Client().
		ApplyURI(dsn).
		SetMaxPoolSize(10).
		SetMinPoolSize(2).
		SetMaxConnIdleTime(5 * time.Minute)
	client, err := mongo.Connect(clientOpts)
	if err != nil {
		return nil, fmt.Errorf("%s: connect mongo: %w", op, err)
	}

	//v2 mongo connect already done ping
	if err := client.Ping(ctx, nil); err != nil {
		_ = client.Disconnect(ctx)
		return nil, fmt.Errorf("%s: ping mongo: %w", op, err)
	}

	return &DB{Client: client}, nil
}

func (db *DB) Close(ctx context.Context) error {
	if db.Client != nil {
		if err := db.Client.Disconnect(ctx); err != nil {
			return err
		}
	}
	return nil
}
