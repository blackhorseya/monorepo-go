package mongodb

import (
	"github.com/blackhorseya/monorepo-go/internal/pkg/configx"
	"github.com/blackhorseya/monorepo-go/pkg/contextx"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// NewClient returns a new mongo client.
func NewClient(config *configx.Config) (*mongo.Client, error) {
	opts := options.Client().ApplyURI(config.Storage.DSN)

	client, err := mongo.Connect(contextx.Background(), opts)
	if err != nil {
		return nil, errors.Wrap(err, "connect mongodb")
	}

	return client, nil
}
