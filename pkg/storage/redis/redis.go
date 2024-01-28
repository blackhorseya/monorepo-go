package redis

import (
	"github.com/blackhorseya/monorepo-go/pkg/configx"
	"github.com/blackhorseya/monorepo-go/pkg/contextx"
	"github.com/redis/go-redis/v9"
	"github.com/testcontainers/testcontainers-go"
	redisc "github.com/testcontainers/testcontainers-go/modules/redis"
)

// Client is a redis client.
type Client struct {
	*redis.Client
}

// NewClient creates a redis client.
func NewClient() (*Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     configx.A.Storage.Redis.Addr,
		Password: configx.A.Storage.Redis.Password,
		DB:       configx.A.Storage.Redis.DB,
	})

	return &Client{
		Client: client,
	}, nil
}

// Container represents the mongodb container type used in the module.
type Container struct {
	*redisc.RedisContainer
}

// NewContainer creates an instance of the mongodb container type.
func NewContainer(ctx contextx.Contextx) (*Container, error) {
	container, err := redisc.RunContainer(
		ctx,
		testcontainers.WithImage("redis:7"),
		redisc.WithLogLevel(redisc.LogLevelVerbose),
	)
	if err != nil {
		return nil, err
	}

	return &Container{container}, nil
}
