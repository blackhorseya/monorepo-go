package mongodb

import (
	"time"

	"github.com/blackhorseya/monorepo-go/entity/orianna/domain/market/agg"
)

type stock struct {
	Symbol           string    `bson:"_id"`
	Name             string    `bson:"name"`
	IndustryCategory string    `bson:"industry_category"`
	ExchangeName     string    `bson:"exchange_name"`
	CreatedAt        time.Time `bson:"created_at"`
	UpdatedAt        time.Time `bson:"updated_at"`
}

func fromStock(v agg.Stock) stock {
	now := time.Now()

	return stock{
		Symbol:           v.GetSymbol(),
		Name:             v.GetName(),
		IndustryCategory: v.GetIndustryCategory(),
		ExchangeName:     v.GetExchangeName(),
		CreatedAt:        now,
		UpdatedAt:        now,
	}
}
