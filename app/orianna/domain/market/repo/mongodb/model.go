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

type stockQuota struct {
	Open      float64   `bson:"open"`
	High      float64   `bson:"high"`
	Close     float64   `bson:"close"`
	Low       float64   `bson:"low"`
	Volume    int       `bson:"volume"`
	Value     float64   `bson:"value"`
	Timestamp time.Time `bson:"timestamp"`
}

func fromStockQuota(v model.StockQuota) stockQuota {
	return stockQuota{
		Open:      v.GetOpen(),
		High:      v.GetHigh(),
		Close:     v.GetClose(),
		Low:       v.GetLow(),
		Volume:    v.GetVolume(),
		Value:     v.GetValue(),
		Timestamp: v.GetUpdatedAt(),
	}
}
