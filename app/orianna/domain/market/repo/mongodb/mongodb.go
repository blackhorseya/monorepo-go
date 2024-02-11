package mongodb

import (
	"github.com/blackhorseya/monorepo-go/entity/orianna/domain/market/agg"
	"github.com/blackhorseya/monorepo-go/entity/orianna/domain/market/repo"
	"github.com/blackhorseya/monorepo-go/pkg/contextx"
	"go.mongodb.org/mongo-driver/mongo"
)

type impl struct {
	client *mongo.Client
}

// NewStockRepo is the constructor of stock repository.
func NewStockRepo(client *mongo.Client) (repo.IStockRepo, error) {
	return &impl{client: client}, nil
}

func (i *impl) List(ctx contextx.Contextx) ([]agg.Stock, error) {
	// TODO implement me
	panic("implement me")
}

func (i *impl) Add(ctx contextx.Contextx, stock agg.Stock) error {
	// TODO implement me
	panic("implement me")
}
