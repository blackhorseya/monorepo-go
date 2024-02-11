//go:build wireinject

//go:generate wire

package biz

import (
	"github.com/blackhorseya/monorepo-go/app/orianna/domain/market/repo/mongodb"
	marketB "github.com/blackhorseya/monorepo-go/entity/orianna/domain/market/biz"
	mongodbx "github.com/blackhorseya/monorepo-go/pkg/storage/mongodb"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	NewMarketBiz,
	mongodb.NewStockRepo,
)

func newForTest() (marketB.IMarketBiz, error) {
	panic(wire.Build(
		NewMarketBiz,
		mongodb.NewStockRepo,
		mongodbx.NewClient,
	))
}
