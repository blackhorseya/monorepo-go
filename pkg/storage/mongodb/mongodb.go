package mongodb

import (
	"github.com/blackhorseya/monorepo-go/pkg/configx"
	"github.com/blackhorseya/monorepo-go/pkg/contextx"
	"github.com/pkg/errors"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/mongodb"
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

// Container is used to represent a mongodb container.
type Container struct {
	*mongodb.MongoDBContainer
}

// NewContainer returns a new mongodb container.
func NewContainer(ctx contextx.Contextx) (*Container, error) {
	container, err := mongodb.RunContainer(ctx, testcontainers.WithImage("mongo:6"))
	if err != nil {
		return nil, errors.Wrap(err, "run mongodb container")
	}

	return &Container{
		MongoDBContainer: container,
	}, nil
}
