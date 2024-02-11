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

func (x *Stock) GetName() string {
	return x.stock.Name
}

func (x *Stock) GetIndustryCategory() string {
	return x.stock.IndustryCategory
}

func (x *Stock) GetExchangeName() string {
	return x.stock.ExchangeName
}
