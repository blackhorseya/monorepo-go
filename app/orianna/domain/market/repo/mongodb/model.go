package mongodb

import (
	"time"

	"github.com/blackhorseya/monorepo-go/entity/orianna/domain/market/agg"
	"github.com/blackhorseya/monorepo-go/entity/orianna/domain/market/model"
)

type stock struct {
	Symbol           string    `bson:"_id"`
	Name             string    `bson:"name"`
	IndustryCategory string    `bson:"industry_category"`
	ExchangeName     string    `bson:"exchange_name"`
	UpdatedAt        time.Time `bson:"updated_at"`
}

// ToAggregate is to convert to aggregate stock.
func (x *stock) ToAggregate() agg.Stock {
	return agg.NewStock(&model.Stock{
		Symbol:           x.Symbol,
		Name:             x.Name,
		IndustryCategory: x.IndustryCategory,
		ExchangeName:     x.ExchangeName,
	})
}
