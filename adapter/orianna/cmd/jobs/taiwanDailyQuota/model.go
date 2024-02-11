package main

import (
	"strconv"

	"github.com/blackhorseya/monorepo-go/entity/domain/market/model"
)

// StockDayResponse is the response of the stock day.
type StockDayResponse struct {
	Code         string `json:"Code"`
	Name         string `json:"Name"`
	TradeVolume  string `json:"TradeVolume"`
	TradeValue   string `json:"TradeValue"`
	OpeningPrice string `json:"OpeningPrice"`
	HighestPrice string `json:"HighestPrice"`
	LowestPrice  string `json:"LowestPrice"`
	ClosingPrice string `json:"ClosingPrice"`
	Change       string `json:"Change"`
	Transaction  string `json:"Transaction"`
}

// ToEntity convert to entity.
func (s *StockDayResponse) ToEntity() *model.Stock {
	price, _ := strconv.ParseFloat(s.ClosingPrice, 64)

	return &model.Stock{
		Symbol: s.Code,
		Name:   s.Name,
		Price:  price,
	}
}
