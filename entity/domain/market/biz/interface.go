//go:generate mockgen -destination=./mock_${GOFILE} -package=${GOPACKAGE} -source=${GOFILE}

package biz

import (
	"github.com/blackhorseya/monorepo-go/entity/domain/market/model"
	"github.com/blackhorseya/monorepo-go/pkg/contextx"
)

// ListStocksOptions is the options for list stocks.
type ListStocksOptions struct {
	Page int
	Size int
}

// IMarketBiz is the interface for market biz.
type IMarketBiz interface {
	// ListStocks is used to list stocks.
	ListStocks(ctx contextx.Contextx, options ListStocksOptions) (stocks []*model.StockInfo, total int, err error)

	// GetStockBySymbol is used to get stock by symbol.
	GetStockBySymbol(ctx contextx.Contextx, symbol string) (stock *model.Stock, err error)
}
