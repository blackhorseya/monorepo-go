package main

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
