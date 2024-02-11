//go:build wireinject

//go:generate wire

package mongodb

import (
	"github.com/blackhorseya/monorepo-go/entity/orianna/domain/market/repo"
	mongodbx "github.com/blackhorseya/monorepo-go/pkg/storage/mongodb"
	"github.com/google/wire"
)

func newForTest() (repo.IStockRepo, error) {
	panic(wire.Build(
		mongodbx.NewClient,
		NewStockRepo,
	))
}
