package agg

import (
	"github.com/blackhorseya/monorepo-go/entity/orianna/domain/market/model"
)

// Stock is an aggregate root that represents the stock.
type Stock struct {
	stock       *model.Stock
	recentQuota model.StockQuota
}

func (x *Stock) GetSymbol() string {
	return x.stock.Symbol
}

func (x *Stock) SetSymbol(symbol string) {
	if x.stock == nil {
		x.stock = &model.Stock{}
	}

	x.stock.Symbol = symbol
}
