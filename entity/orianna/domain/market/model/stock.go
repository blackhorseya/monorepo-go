package model

// Stock is an entity that represents the stock.
type Stock struct {
	Symbol           string `json:"symbol"`
	Name             string `json:"name,omitempty"`
	IndustryCategory string `json:"industry_category,omitempty"`
	ExchangeName     string `json:"exchange_name,omitempty"`
}
