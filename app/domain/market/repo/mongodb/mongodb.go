package mongodb

import (
	"github.com/blackhorseya/monorepo-go/app/domain/market/repo"
	"github.com/blackhorseya/monorepo-go/entity/domain/market/model"
	"github.com/blackhorseya/monorepo-go/pkg/contextx"
	"go.mongodb.org/mongo-driver/mongo"
)

type impl struct {
	rw *mongo.Client
}

// NewStorager is the factory method to create a Storager instance.
func NewStorager(rw *mongo.Client) (repo.Storager, error) {
	return &impl{rw: rw}, nil
}

func (i *impl) GetBySymbol(ctx contextx.Contextx, symbol string) (info *model.StockInfo, err error) {
	// todo: 2024/2/5|sean|implement me
	panic("implement me")
}
