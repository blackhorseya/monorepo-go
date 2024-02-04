package mongodb

import (
	"github.com/blackhorseya/monorepo-go/pkg/configx"
	"github.com/blackhorseya/monorepo-go/pkg/contextx"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// NewClient returns a new mongo client.
func NewClient() (*mongo.Client, error) {
	opts := options.Client().ApplyURI(configx.A.Storage.Mongodb.DSN)

	client, err := mongo.Connect(contextx.Background(), opts)
	if err != nil {
		return nil, errors.Wrap(err, "connect mongodb")
	}

	return client, nil
}
