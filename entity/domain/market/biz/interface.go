//go:generate mockgen -destination=./mock_${GOFILE} -package=${GOPACKAGE} -source=${GOFILE}

package biz

import (
	"github.com/blackhorseya/monorepo-go/entity/domain/market/model"
	"github.com/blackhorseya/monorepo-go/pkg/contextx"
)

// IMarketBiz is the interface for market biz.
type IMarketBiz interface {
	// GetStockBySymbol is used to get stock by symbol.
	GetStockBySymbol(ctx contextx.Contextx, symbol string) (stock *model.Stock, err error)
}
