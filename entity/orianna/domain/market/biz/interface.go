//go:generate mockgen -destination=./mock_${GOFILE} -package=${GOPACKAGE} -source=${GOFILE}

package biz

import (
	"github.com/blackhorseya/monorepo-go/entity/orianna/domain/market/agg"
	"github.com/blackhorseya/monorepo-go/pkg/contextx"
)

// IMarketBiz is the interface for market biz.
type IMarketBiz interface {
	ListStocks(ctx contextx.Contextx) ([]agg.Stock, error)
	GetStockBySymbol(ctx contextx.Contextx, symbol string) (agg.Stock, error)
}
