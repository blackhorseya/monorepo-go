package agg

import (
	"encoding/json"

	"github.com/blackhorseya/monorepo-go/entity/orianna/domain/market/model"
)

// Stock is an aggregate root that represents the stock.
type Stock struct {
	stock       *model.Stock
	recentQuota model.StockQuota
}

// NewStock is the constructor of Stock.
func NewStock(stock *model.Stock) Stock {
	return Stock{
		stock: stock,
	}
}

func (x *Stock) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		*model.Stock
	}{
		Stock: x.stock,
	})
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
