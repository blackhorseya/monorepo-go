//go:generate mockgen -destination=./mock_${GOFILE} -package=${GOPACKAGE} -source=${GOFILE}

package biz

import (
	"time"

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
	// GetMarketInfoByType is used to get market info.
	GetMarketInfoByType(ctx contextx.Contextx, typeStr string, t *time.Time) (info *model.MarketInfo, err error)

	// GetStockBySymbol is used to get stock by symbol.
	GetStockBySymbol(ctx contextx.Contextx, symbol string) (stock *model.Stock, err error)

	// ListStocks is used to list stocks.
	ListStocks(ctx contextx.Contextx, options ListStocksOptions) (stocks []*model.StockInfo, total int, err error)
}
